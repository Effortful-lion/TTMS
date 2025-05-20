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
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port), // 连接地址和端口
		Password: cfg.Password,
		DB:       cfg.Db,
	})
	_, err = Rdb.Ping(context.Background()).Result()  // 添加 context 参数
	if err == nil {
		log.Println("redis连接成功！")
		return nil	
	}
    return err
}