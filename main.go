package main

import (
	"TTMS/api"
	"TTMS/config"
	"TTMS/dao/mysql"
	"TTMS/dao/redis"
	_ "TTMS/docs"
	"fmt"
	"sync"
)

// @title 后端系统 API在线测试文档
// @version 1.0
// @description 这是一个简单的后端系统 API 文档，包含用户管理、视频管理等功能。
// @termsOfService http://example.com/terms/
// @contact.name Server-lion
// @contact.url https://github.com/Effortful-lion
// @contact.email server-lion@qq.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host 45.95.212.18:43223
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

	// 初始化路由
	r := api.InitRouter()

	var wg sync.WaitGroup
	wg.Add(2)

	// http 服务
	go func ()  {
		defer wg.Done()
		port := config.Conf.AppConfig.HttpPort
	    if err := r.Run(fmt.Sprintf(":%d", port)); err!= nil {
			fmt.Printf("run server failed:%v\n",err)
			return
		}	
	}()

	// https 服务
	go func ()  {
		defer wg.Done()
		port := config.Conf.AppConfig.HttpsPort
	    if err := r.RunTLS(fmt.Sprintf(":%d",port), "./https/server.pem", "./https/server.key"); err!= nil {
			fmt.Printf("run server failed:%v\n",err)
			return
	    }	
	}()

	wg.Wait()
}