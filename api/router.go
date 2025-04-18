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

	// 普通用户权限模块
	userGroup := r.Group("/user")
	userGroup.Use(middleware.JWTAuthMiddleware())
	{
		// test 路由
		userGroup.GET("/auth", func(ctx *gin.Context) {
			auth := controller.GetCurrentUserAuthority(ctx)
			ctx.JSON(200, gin.H{"auth": auth})
		})

	}
	
	// 管理员权限模块
	adminGroup := r.Group("/admin")
	adminGroup.Use(middleware.JWTAuthMiddleware(), middleware.AdminAuthMiddleware())
	{
		// test 路由
		adminGroup.GET("/auth", func(ctx *gin.Context) {
			auth := controller.GetCurrentUserAuthority(ctx)
			ctx.JSON(200, gin.H{"auth": auth})
		})

		// 演出厅增删改查
		adminGroup.POST("/hall", controller.NewHallController().AddHallHandler)
		adminGroup.DELETE("/hall/:hall_id", controller.NewHallController().DeleteHallHandler)
		adminGroup.PUT("/hall", controller.NewHallController().UpdateHallHandler)
		adminGroup.GET("/hall", controller.NewHallController().GetHallListHandler)
		adminGroup.GET("/hall/:hall_id", controller.NewHallController().GetHallHandler)
	}

	// 运营经理权限模块
	managerGroup := r.Group("/manager")
	managerGroup.Use(middleware.JWTAuthMiddleware(), middleware.ManagerAuthMiddleware())              
	{
		// test 路由
		managerGroup.GET("/auth", func(ctx *gin.Context) {
			auth := controller.GetCurrentUserAuthority(ctx)
			ctx.JSON(200, gin.H{"auth": auth})
		})
		// 剧目增删改查
		managerGroup.POST("/play", controller.NewPlayController().AddPlayHandler)
		managerGroup.DELETE("/play/:play_id", controller.NewPlayController().DeletePlayHandler)
		managerGroup.PUT("/play", controller.NewPlayController().UpdatePlayHandler)
		managerGroup.GET("/play", controller.NewPlayController().GetPlayListHandler)
		managerGroup.GET("/play/:play_id", controller.NewPlayController().GetPlayHandler)
	}



	return r
}