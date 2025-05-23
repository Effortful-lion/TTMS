package middleware

import (
	"TTMS/controller"
	"TTMS/pkg"
	"TTMS/pkg/resp"
	"net/http"
	"strings"
	"TTMS/dao/redis"

	"github.com/gin-gonic/gin"
)

// 拦截过期token 和 无权限请求的中间件
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 这里假设Token放在Header的Authorization中，并使用Bearer开头
		// 这里的具体实现方式要依据你的实际业务情况决定
		authHeader := c.Request.Header.Get("Authorization")
		// Token判空
		if authHeader == "" {
			c.JSON(http.StatusOK, gin.H{
				"code": resp.CodeNeedLogin,
				"msg":  "请求头中token为空",
			})
			c.Abort()
			return
		}
		// Token格式判断
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusOK, gin.H{
				"code": resp.CodeNeedLogin,
				"msg":  "请求头中token格式有误",
			})
			c.Abort()
			return
		}
		// Token解析
		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		mc, err := pkg.ParseToken(parts[1])
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": resp.CodeNeedLogin,
				"msg":  "Token解析错误: " + err.Error(),
			})
			c.Abort()
			return
		}

		c.Set(controller.ContextUserIDKey, int64(mc.UserID))
		c.Set(controller.ContextUserAuthorityKey, mc.Authority)
		
		// 检查用户权限
		url := c.Request.URL.Path
		resource := strings.Split(url, "/")[1]
		// 检查用户角色权限是否足够
		auth := mc.Authority
		if redis.SetResourceIsMember(resource, auth) {
			c.Next()
		} else {
			c.AbortWithStatusJSON(200, gin.H{"code": 401, "msg": "权限不足"})
		}
	}
}