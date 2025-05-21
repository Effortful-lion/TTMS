package redis

import (
	"TTMS/config"
	"context" // 新增 context 包
	"fmt"
	"log"

	"github.com/redis/go-redis/v9" // 修改导入路径
)

// redis数据库连接

var Rdb *redis.Client

// 初始化连接
func InitRedis() (err error) {
    cfg := config.Conf.RedisConfig
    Rdb = redis.NewClient(&redis.Options{
        Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
        Password: cfg.Password,
        DB:       cfg.Db,
    })

    ctx := context.Background()
    
    // 测试连接
    _, err = Rdb.Ping(ctx).Result()
    if err != nil {
        log.Printf("Redis ping failed: %v", err)
        return err // 连接失败时直接返回错误
    }
    log.Println("redis连接成功！")

    return nil
}