package middleware

import (
	"TTMS/controller"
	"TTMS/pkg"
	"TTMS/pkg/resp"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// JWTAuthMiddleware 基于JWT的认证中间件
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 这里假设Token放在Header的Authorization中，并使用Bearer开头
		// 这里的具体实现方式要依据你的实际业务情况决定
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusOK, gin.H{
				"code": resp.CodeNeedLogin,
				"msg":  "请求头中token为空",
			})
			c.Abort()
			return
		}
		// 按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusOK, gin.H{
				"code": resp.CodeNeedLogin,
				"msg":  "请求头中token格式有误",
			})
			c.Abort()
			return
		}
		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		mc, err := pkg.ParseToken(parts[1])
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": resp.CodeNeedLogin,
				"msg":  "Token过期: " + err.Error(),
			})
			c.Abort()
			return
		}

		// token解析到claims后
		// 将当前请求的userID信息保存到请求的上下文context c中：
		c.Set(controller.ContextUserIDKey, int64(mc.UserID))
		// 将当前请求的权限信息保存到上下文中
		c.Set(controller.ContextUserAuthorityKey, mc.Authority)
		// 后续的处理请求的函数可以用c.Get("userId")来获取当前请求的用户信息
		c.Next()
	}
}

func JWTNoAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 这里假设Token放在Header的Authorization中，并使用Bearer开头
		// 这里的具体实现方式要依据你的实际业务情况决定
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.Next()
			return
		}
		// 按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusOK, gin.H{
				"code": resp.CodeNeedLogin,
				"msg":  "请求头中token格式有误",
			})
			c.Abort()
			return
		}
		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		mc, err := pkg.ParseToken(parts[1])
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": resp.CodeNeedLogin,
				"msg":  "Token过期: " + err.Error(),
			})
			c.Abort()
			return
		}

		// token解析到claims后
		// 将当前请求的userID信息保存到请求的上下文context c中：
		c.Set(controller.ContextUserIDKey, int64(mc.UserID))
		c.Set(controller.ContextUserAuthorityKey, mc.Authority)
		// 后续的处理请求的函数可以用c.Get("userId")来获取当前请求的用户信息
		c.Next()
	}
}

func AdminAndManagerAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		auth := c.GetString(controller.ContextUserAuthorityKey)
		if auth!= pkg.AuthAdmin && auth!= pkg.AuthManager {
			resp.ResponseError(c, resp.CodeNoAuth)
			c.Abort()
		}
		c.Next()
	}
}

func AdminAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		auth := c.GetString(controller.ContextUserAuthorityKey)
		if auth != pkg.AuthAdmin {
			resp.ResponseError(c, resp.CodeNoAuth)
			c.Abort()
		}
		c.Next()
	}	
}

func UserAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		auth := c.GetString(controller.ContextUserAuthorityKey)
		if auth != pkg.AuthUser {
			resp.ResponseError(c, resp.CodeNoAuth)
			c.Abort()
		}
		c.Next()
	}	
}

func StaffAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		auth := c.GetString(controller.ContextUserAuthorityKey)
		if auth != pkg.AuthStaff {
			resp.ResponseError(c, resp.CodeNoAuth)
			c.Abort()
		}
		c.Next()
	}	
}

func ManagerAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		auth := c.GetString(controller.ContextUserAuthorityKey)
		if auth != pkg.AuthManager {
			resp.ResponseError(c, resp.CodeNoAuth)
			c.Abort()
		}
		c.Next()
	}	
}

func FinanceAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		auth := c.GetString(controller.ContextUserAuthorityKey)
		if auth != pkg.AuthFinance {
			resp.ResponseError(c, resp.CodeNoAuth)
			c.Abort()
		}
		c.Next()
	}	
}

func TicketorAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		auth := c.GetString(controller.ContextUserAuthorityKey)
		if auth != pkg.AuthTicketor {
			resp.ResponseError(c, resp.CodeNoAuth)
			c.Abort()
		}
		c.Next()
	}	
}

func AccountAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		auth := c.GetString(controller.ContextUserAuthorityKey)
		if auth != pkg.AuthAccount {
			resp.ResponseError(c, resp.CodeNoAuth)
			c.Abort()
		}
		c.Next()
	}	
}
