package redis

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"

	"TTMS/dao/mysql"
	"TTMS/model/do"
	"TTMS/pkg/common"

	"github.com/redis/go-redis/v9"
)

const (
	PlanStatusKeyPrefix  = "plan_status:"
	PlanEndTimeKeyPrefix = "plan_end_time:"
	SyncInterval         = 30 * time.Second
	FinalExpiration      = 1 * time.Minute
	EndTimeExpire		 = 1 * time.Hour
)

type RedisPlanManager struct {
	client *redis.Client
	wg     *sync.WaitGroup
	ctx    context.Context
	cancel context.CancelFunc
	mysqlcli *mysql.PlanDao
}

var RedisPlanCli *RedisPlanManager

func NewRedisPlanManager(client *redis.Client) *RedisPlanManager {
	ctx, cancel := context.WithCancel(context.Background())
	return &RedisPlanManager{
		client: client,
		wg:     &sync.WaitGroup{},
		ctx:    ctx,
		cancel: cancel,
		mysqlcli: mysql.NewPlanDao(),
	}
}

// 启动所有服务
func (r *RedisPlanManager) Start() {
	r.wg.Add(3) // 监听、更新、同步三个goroutine
	go r.watchPlanStatus()
	go r.updatePlanStatus()
	go r.syncPlanStatus()
	fmt.Println("RedisPlanManager 启动成功")
}

// 优雅关闭所有服务
func (r *RedisPlanManager) Close() {
	r.cancel()    // 取消上下文
	r.wg.Wait()   // 等待所有goroutine完成
}

// 设置演出计划状态为"未开始"
func (r *RedisPlanManager) SetPlanStatusBefore(planId int64, startTime, endTime string) error {
	startTimeStamp := common.ParseStringTimeToTimeStamp(startTime)
	now := time.Now().Unix()
	expire := startTimeStamp - now

	if expire <= 0 {
		return errors.New("演出开始时间已过")
	}

	// 事务化设置状态和结束时间
	tx := r.client.TxPipeline()
	key := fmt.Sprintf(PlanStatusKeyPrefix+"%d", planId)
	endTimeKey := fmt.Sprintf(PlanEndTimeKeyPrefix+"%d", planId)
	
	tx.Set(r.ctx, key, int(do.PlanStatusBefore), time.Duration(expire)*time.Second)
	tx.Set(r.ctx, endTimeKey, endTime, EndTimeExpire)
	
	_, err := tx.Exec(r.ctx)
	return err
}

// 监听过期事件
func (r *RedisPlanManager) watchPlanStatus() {
	defer r.wg.Done()
	
	// 确保键空间通知启用
	config, err := r.client.ConfigGet(r.ctx, "notify-keyspace-events").Result()
	if err != nil || config["notify-keyspace-events"] != "Ex" {
		if err := r.client.ConfigSet(r.ctx, "notify-keyspace-events", "Ex").Err(); err != nil {
			fmt.Printf("启用键空间通知失败: %v\n", err)
			return
		}
	}

	// 订阅过期事件
	pubsub := r.client.PSubscribe(r.ctx, "__keyevent@0__:expired")
	defer pubsub.Close()

	// Lua脚本：原子化获取键值（避免键过期后无法读取）
	luaScript := `
		local status = redis.call('GET', KEYS[1])
		if status then
			return status
		end
		return nil
	`

	for {
		select {
		case msg, ok := <-pubsub.Channel():
			if !ok {
				return // 通道关闭
			}
			
			key := msg.Payload
			if !strings.HasPrefix(key, PlanStatusKeyPrefix) {
				continue
			}

			planIdStr := strings.TrimPrefix(key, PlanStatusKeyPrefix)
			if planIdStr == "" {
				continue
			}

			// 执行Lua脚本获取状态
			status, err := r.client.Eval(r.ctx, luaScript, []string{key}).Result()
			if err != nil && err != redis.Nil {
				fmt.Printf("获取状态失败: %v\n", err)
				continue
			}

			var statusInt int
			if status != nil {
				statusInt, _ = strconv.Atoi(status.(string))
			}

			// 根据状态值发送通知
			switch statusInt {
			case int(do.PlanStatusBefore):
				fmt.Printf("通知演出开始: planID=%s\n", planIdStr)
				go func(id string) {
					select {
					case NotifyStart <- id:
					case <-r.ctx.Done():
					}
				}(planIdStr)
			case int(do.PlanStatusDuring):
				fmt.Printf("通知演出结束: planID=%s\n", planIdStr)
				go func(id string) {
					select {
					case NotifyEnd <- id:
					case <-r.ctx.Done():
					}
				}(planIdStr)
			}

		case <-r.ctx.Done():
			return
		}
	}
}

// 更新计划状态
func (r *RedisPlanManager) updatePlanStatus() {
	defer r.wg.Done()
	
	for {
		select {
		case planIdStr := <-NotifyStart:
			planId, err := strconv.ParseInt(planIdStr, 10, 64)
			if err != nil {
				continue
			}
			
			// 获取结束时间
			endTimeKey := fmt.Sprintf(PlanEndTimeKeyPrefix+"%d", planId)
			endTime, err := r.client.Get(r.ctx, endTimeKey).Result()
			if err != nil {
				fmt.Printf("获取结束时间失败: %v\n", err)
				continue
			}

			// 计算剩余时间
			endTimeStamp := common.ParseStringTimeToTimeStamp(endTime)
			now := time.Now().Unix()
			expire := endTimeStamp - now

			// 更新状态
			key := fmt.Sprintf(PlanStatusKeyPrefix+"%d", planId)
			if expire <= 0 {
				// 已结束
				r.client.Set(r.ctx, key, int(do.PlanStatusAfter), FinalExpiration)
				r.client.Expire(r.ctx, endTimeKey, FinalExpiration)
				fmt.Printf("演出已结束: planID=%s\n", planIdStr)
			} else {
				// 进行中
				r.client.Set(r.ctx, key, int(do.PlanStatusDuring), time.Duration(expire)*time.Second)
				fmt.Printf("演出进行中: planID=%s，剩余时间=%d秒\n", planIdStr, expire)
			}

		case planIdStr := <-NotifyEnd:
			planId, err := strconv.ParseInt(planIdStr, 10, 64)
			if err != nil {
				continue
			}
			
			// 设置为已结束状态
			key := fmt.Sprintf(PlanStatusKeyPrefix+"%d", planId)
			r.client.Set(r.ctx, key, int(do.PlanStatusAfter), FinalExpiration)
			
			// 设置结束时间键的过期时间
			endTimeKey := fmt.Sprintf(PlanEndTimeKeyPrefix+"%d", planId)
			r.client.Expire(r.ctx, endTimeKey, FinalExpiration)
			
			fmt.Printf("演出已结束: planID=%s\n", planIdStr)

		case <-r.ctx.Done():
			return
		}
	}
}

// 同步状态到数据库
func (r *RedisPlanManager) syncPlanStatus() {
	defer r.wg.Done()
	
	ticker := time.NewTicker(SyncInterval)
	defer ticker.Stop()
	
	for {
		select {
		case <-ticker.C:
			keys, err := r.scanKeys(PlanStatusKeyPrefix + "*")
			if err != nil {
				fmt.Printf("扫描键失败: %v\n", err)
				continue
			}

			for _, key := range keys {
				planIdStr := strings.TrimPrefix(key, PlanStatusKeyPrefix)
				status, err := r.client.Get(r.ctx, key).Int()
				if err != nil {
					continue
				}
				
				if err := r.updatePlanStatusInDB(planIdStr, status); err != nil {
					fmt.Printf("同步失败: planID=%s, err=%v\n", planIdStr, err)
				}
			}

		case <-r.ctx.Done():
			return
		}
	}
}

// 辅助函数：扫描匹配的键
func (r *RedisPlanManager) scanKeys(pattern string) ([]string, error) {
	var keys []string
	iter := r.client.Scan(r.ctx, 0, pattern, 0).Iterator()
	
	for iter.Next(r.ctx) {
		keys = append(keys, iter.Val())
	}
	
	if err := iter.Err(); err != nil {
		return nil, err
	}
	
	return keys, nil
}

// 更新数据库中的计划状态（示例实现）
func (r *RedisPlanManager) updatePlanStatusInDB(planIdStr string, status int) error {
	planId, err := strconv.ParseInt(planIdStr, 10, 64)
	if err != nil {
		return err
	}
	
	// fmt.Printf("同步到数据库: planID=%d, status=%d\n", planId, status)
	err = r.mysqlcli.UpdatePlanStatus(planId, int8(status))
	if err != nil {
		return err
	}
	return nil
}

func (r *RedisPlanManager) DeletePlanStatus(planId int64) error {
	// 删除 以 plan_status:planID 为 key 的键
	// 删除 以 plan_end_time:planID 为 key 的键
	err := r.client.Del(r.ctx, PlanStatusKeyPrefix+strconv.FormatInt(planId, 10)).Err()
	if err != nil {
		return err
	}
	err = r.client.Del(r.ctx, PlanEndTimeKeyPrefix+strconv.FormatInt(planId, 10)).Err()
	if err != nil {
		return err
	}
	return nil
}

// 全局通知通道
var (
	NotifyStart = make(chan string, 100)
	NotifyEnd   = make(chan string, 100)
)