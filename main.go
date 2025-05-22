package main

import (
	"TTMS/api"
	"TTMS/config"
	"TTMS/dao/mysql"
	"TTMS/dao/redis"
	"fmt"
)

func main() {

	// 初始化配置文件
	if err := config.InitConfig(); err != nil {
		fmt.Println(err)
	}

	// 初始化数据库
	if err := mysql.InitMysql(); err != nil {
		fmt.Println(err)
	}

	// 初始化redis
	if err := redis.InitRedis(); err != nil {
		fmt.Println(err)
	}

	// 初始化redis存储的资源访问权限表
	if err := redis.InitRedisResource(); err!= nil {
		fmt.Println(err)
	}

	// 启动状态监听
	go redis.RedisPlanCli.Start()

	r := api.InitRouter()

	httpPort := config.Conf.AppConfig.HttpPort
	r.Run(fmt.Sprintf(":%d", httpPort))

}
