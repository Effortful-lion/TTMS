package redis

import (
	"context"
	"fmt"
	"time"
)

// 演出状态

type PlanRedis struct {
}

func NewPlanRedis() *PlanRedis {
	return &PlanRedis{}
}

func (*PlanRedis) SetPlanStatus(plan_id int64, status int, expire int64) error {
	err := Rdb.Set(context.Background(), fmt.Sprintf("plan:%d", plan_id), status, time.Duration(expire)*time.Second).Err()
	if err!= nil {
		return err
	}
	return nil
}

func (*PlanRedis) GetPlanStatus(plan_id int64) (int, error) {
	status, err := Rdb.Get(context.Background(), fmt.Sprintf("plan:%d", plan_id)).Int()
	if err != nil {
		return 0, err
	}
	return status, nil	
}

// 检查是否演出结束
func (*PlanRedis) CheckPlanStatus(plan_id int64) (bool, error) {
	status, err := Rdb.Get(context.Background(), fmt.Sprintf("plan:%d", plan_id)).Int()
	if err!= nil {
		return false, err
	}
	return (status == 2 && status != 0), nil
}

// TODO 定时任务更新数据库演出状态