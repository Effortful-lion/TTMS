package main

import (
	"TTMS/api"
	"TTMS/config"
	"TTMS/dao/mysql"
	"TTMS/dao/redis"
	_ "TTMS/docs"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"log"
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
// @host 39.105.136.3:8888
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

	srv := &http.Server{
		Addr:    fmt.Sprintf("localhost:%d", config.Conf.AppConfig.Port),
		Handler: r,
	}

	// 优雅关闭
	go func(*gin.Engine, *http.Server) {
		log.Printf("server listening addr: %s\n", srv.Addr)
		if err := srv.ListenAndServe(); err!= nil && err!= http.ErrServerClosed {
			fmt.Println(err)
		}
	}(r, srv)
	quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit
		// 等待两秒，确保所有请求都处理完
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		if err := srv.Shutdown(ctx); err!= nil {
			fmt.Println(err)
		}
		// cancel() 函数用于主动取消通过`context.WithTimeout` 创建的上下文及其关联资源
		defer cancel()
		fmt.Println("Shutdown Server...")
}