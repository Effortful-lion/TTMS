package resp

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 定义响应信息结构体

// 响应结构体
type ResponseData struct {
	Status ResCode    `json:"status"` // 自定义的status
	Msg  any   `json:"msg"`  // 自定义的msg
	Data any 	`json:"data"`// 自定义的数据  ,omitempty可忽略空值不展示  Data any 	`json:"data,omitempty"`
}


// 响应错误信息：code+错误信息
func ResponseError(c *gin.Context, status ResCode){
	c.JSON(http.StatusOK,&ResponseData{
		Status: status,
		Msg:status.Msg(),
		Data:nil,
	})
}

// 响应成功信息：
func ResponseSuccess(c *gin.Context, data any){
	c.JSON(http.StatusOK,&ResponseData{
		Status: CodeSuccess,
		Msg: CodeSuccess.Msg(),
		Data: data,
	})
}

// 响应具体错误信息
func ResponseErrorWithMsg(c *gin.Context, status ResCode, msg string){
	c.JSON(http.StatusOK,&ResponseData{
		Status: status,
		Msg: msg,
		Data: nil,
	})
}