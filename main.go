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
	if err := mysql.InitMysql(); err!= nil {
		fmt.Println(err)	
	}

	// 初始化redis
	if err := redis.InitRedis(); err!= nil {
		fmt.Println(err)
	}

	r := api.InitRouter()

	r.Run()

}