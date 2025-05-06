package api

import (
	"TTMS/controller"
	"TTMS/middleware"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/files"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	// 跨域中间件
	r.Use(middleware.Cors())

	//注册swagger api相关路由
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 测试路由
	r.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "hello world",
		})
	})

	r.POST("/signup", controller.NewUserController().SignUpHandler) // 用户注册
	r.POST("/login", controller.NewUserController().LoginHandler)   // 用户登录
	r.POST("/userinfo", middleware.JWTNoAuthMiddleware(),controller.NewUserController().GetUserInfoHandler) // 用户信息

	// manage 路由组
	manageGroup := r.Group("/manage")
	manageGroup.Use(middleware.JWTAuthMiddleware()) // 应用JWT认证中间件

	{
		// 这里的路由是不需要权限的，因为它们是公共的

	}

	{
		// 这里的路由是需要权限的，并且权限不是单一的
		// 剧目增删改查
		manageGroup.POST("/play", middleware.AdminAndManagerAuthMiddleware(),controller.NewPlayController().AddPlayHandler)
		manageGroup.DELETE("/play/:play_id", middleware.AdminAndManagerAuthMiddleware(),controller.NewPlayController().DeletePlayHandler)
		manageGroup.PUT("/play", middleware.AdminAndManagerAuthMiddleware(),controller.NewPlayController().UpdatePlayHandler)
		manageGroup.GET("/play", middleware.AdminAndManagerAuthMiddleware(),controller.NewPlayController().GetPlayListHandler)
		manageGroup.GET("/play/:play_id", middleware.AdminAndManagerAuthMiddleware(),controller.NewPlayController().GetPlayHandler)
		// 演出计划增删改查
		manageGroup.POST("/plan", middleware.AdminAndManagerAuthMiddleware(),controller.NewPlanController().AddPlanHandler)
	}

	userGroup := manageGroup.Group("") // 用户管理路由组
	userGroup.Use(middleware.UserAuthMiddleware()) // 应用用户权限中间件
	{
		

	}

	adminGroup := manageGroup.Group("") // 管理员管理路由组
	adminGroup.Use(middleware.AdminAuthMiddleware()) // 应用管理员权限中间件
	{
		// 演出厅增删改查
		adminGroup.POST("/hall", controller.NewHallController().AddHallHandler)
		adminGroup.DELETE("/hall/:hall_id", controller.NewHallController().DeleteHallHandler)
		adminGroup.PUT("/hall", controller.NewHallController().UpdateHallHandler)
		adminGroup.GET("/hall", controller.NewHallController().GetHallListHandler)
		adminGroup.GET("/hall/:hall_id", controller.NewHallController().GetHallHandler)
	}

	managerGroup := manageGroup.Group("") // 运营经理管理路由组
	managerGroup.Use(middleware.ManagerAuthMiddleware()) // 应用运营经理权限中间件
	{
		
	}

	return r
}