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
	r.POST("/userinfo", middleware.JWTAuthMiddleware(),controller.NewUserController().GetUserInfoHandler)

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
	adminGroup.Use(middleware.JWTAuthMiddleware())
	{
		// test 路由
		adminGroup.GET("/auth", func(ctx *gin.Context) {
			auth := controller.GetCurrentUserAuthority(ctx)
			ctx.JSON(200, gin.H{"auth": auth})
		})
	}

	// 用户模块
	// userGroup := r.Group("/user")
	// {
		
	// }



	return r
}