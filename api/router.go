package api

import (
	"TTMS/alipay"
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

	CollectionGroup := r.Group("/collection")
	CollectionGroup.Use(middleware.JWTAuthMiddleware())
	{
		// 暂时，只有管理员具有 数据统计结果 的查看权限
		// 统计所有票房列表
		CollectionGroup.GET("ticket_count", controller.NewTicketController().CountTicketListHandler)
		// 统计单个剧目票房
		CollectionGroup.GET("ticket_count/:play_id", controller.NewTicketController().CountTicketHandler)
		// 统计单场演出票房
		CollectionGroup.GET("ticket_count/once/:plan_id", controller.NewTicketController().CountOnceTicketHandler)
		// 剧目单场票房占比
		CollectionGroup.GET("ticket_count/percentage/:plan_id", controller.NewTicketController().CountOnceTicketPercentageHandler)
		// 单场上座率统计
		CollectionGroup.GET("seat_count/percentage/once/:plan_id", controller.NewTicketController().CountOnceSeatHandler)
		// 剧目上座率统计
		CollectionGroup.GET("seat_count/percentage/:play_id", controller.NewTicketController().CountSeatHandler)
	}

	SaleGroup := r.Group("/sale")
	SaleGroup.Use(middleware.JWTAuthMiddleware())
	{
		// customer/admin（admin为了调试）
		{
			// 暂时规定：管理员、用户可以买、退票
			SaleGroup.POST("/ticket", controller.NewTicketController().BuyHandler)
			// 退票 : 退票的结构体：直接 ticket_id 还是给一堆 ... 其他的信息
			SaleGroup.DELETE("/ticket", controller.NewTicketController().CancelHandler)
			// 查票 ：查询个人的所有票
			SaleGroup.GET("/ticket", controller.NewTicketController().GetTicketListHandler)
			// 核销票
			SaleGroup.PUT("/ticket/verify", controller.NewTicketController().VerifyHandler)

			// 支付宝支付
			SaleGroup.GET("/alipay", alipay.Pay)
			SaleGroup.GET("/alipay/return", alipay.Callback)
			SaleGroup.GET("/alipay/notify", alipay.Notify)
		}

		// employ
		{
			// TODO 卖票：填充 sale 交易信息
			SaleGroup.POST("/sell", controller.NewSaleController().SellHandler)
		}
	}

	ManageGroup := r.Group("/manage")
	ManageGroup.Use(middleware.JWTAuthMiddleware())
	{
		// 剧目管理增删改查
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
