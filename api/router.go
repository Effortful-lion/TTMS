package api

import (
	"TTMS/config"
	"TTMS/controller"
	"TTMS/middleware"
	"fmt"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	// 初始化路由组
	r := gin.Default()

	// 设置跨域
	r.Use(middleware.Cors())

	// 注册路由
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// 注册路由组
	r.POST("/signup", controller.NewUserController().SignUpHandler) // 用户注册
	r.POST("/login", controller.NewUserController().LoginHandler)   // 用户登录
	r.POST("/userinfo", middleware.JWTAuthMiddleware(),controller.NewUserController().GetUserInfoHandler) // 用户信息

	// 按照资源权限分配路由组
	// 然后按照模块分配路由组

	SaleGroup := r.Group("/sale")
	SaleGroup.Use(middleware.JWTAuthMiddleware())
	{
		// 暂时规定：管理员、售票员可以售票 或者 用户可以买票
		SaleGroup.POST("/ticket", controller.NewTicketController().BuyHandler)
	}

	ManageGroup := r.Group("/manage")
	ManageGroup.Use(middleware.JWTAuthMiddleware())
	{
		// 剧目管理
		PlayGroup := ManageGroup.Group("/play")
		{
			PlayGroup.POST("", controller.NewPlayController().AddPlayHandler)
			PlayGroup.DELETE("/:play_id", controller.NewPlayController().DeletePlayHandler)
			PlayGroup.PUT("", controller.NewPlayController().UpdatePlayHandler)
			PlayGroup.GET("", controller.NewPlayController().GetPlayListHandler)
			PlayGroup.GET("/:play_id", controller.NewPlayController().GetPlayHandler)
		}

		// 演出计划增删改查
		PlanGroup := ManageGroup.Group("/plan")
		{
			PlanGroup.POST("", controller.NewPlanController().AddPlanHandler)	
			PlanGroup.DELETE("/:plan_id", controller.NewPlanController().DeletePlanHandler)
			PlanGroup.PUT("", controller.NewPlanController().UpdatePlanHandler)
			PlanGroup.GET("", controller.NewPlanController().GetPlanListHandler)
			PlanGroup.GET("/:plan_id", controller.NewPlanController().GetPlanHandler)
		}

		// 演出厅增删改查
		HallGroup := ManageGroup.Group("/hall")
		{
			HallGroup.POST("", controller.NewHallController().AddHallHandler)
			HallGroup.DELETE("/:hall_id", controller.NewHallController().DeleteHallHandler)
			HallGroup.PUT("", controller.NewHallController().UpdateHallHandler)
			HallGroup.GET("", controller.NewHallController().GetHallListHandler)
			HallGroup.GET("/:hall_id", controller.NewHallController().GetHallHandler)
		}
	}

	port := config.Conf.AppConfig.HttpPort
	r.Run(fmt.Sprintf(":%d", port))
	return r
}
