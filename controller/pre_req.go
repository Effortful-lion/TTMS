package controller

import (
	"github.com/gin-gonic/gin"
)

var ContextUserIDKey string = "userID"

//var ContextUserAuthorityKey string = "authority"

func GetCurrentUserID(c *gin.Context) int64 {
	userid := c.GetInt64(ContextUserIDKey)
	return userid
}

// func GetCurrentUserAuthority(c *gin.Context) string {
// 	auth  := c.GetString(ContextUserAuthorityKey)
// 	return auth
// }
