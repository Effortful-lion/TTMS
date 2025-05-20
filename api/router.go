package api

import (
	"TTMS/config"
	"TTMS/controller"
	"fmt"
	"TTMS/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine{
	// 初始化路由组
	r := gin.Default()

	// 注册路由
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// 注册路由组
	r.POST("/signup", controller.NewUserController().SignUpHandler) // 用户注册
	r.POST("/login", controller.NewUserController().LoginHandler)   // 用户登录
	r.POST("/userinfo", middleware.(),controller.NewUserController().GetUserInfoHandler) // 用户信息



	port := config.Conf.AppConfig.HttpPort
	r.Run(fmt.Sprintf(":%d", port))
	return r
}