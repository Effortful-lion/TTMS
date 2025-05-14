package controller

import (
	"TTMS/model/dto"
	"TTMS/pkg/resp"
	"TTMS/service"

	"github.com/gin-gonic/gin"
)

type PlanHandler struct {
}

func NewPlanController() *PlanHandler {
	return &PlanHandler{}
}


// @Summary 增加演出
// @Description 增加演出
// @Tags 演出管理
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param object body dto.PlanInsertReq true "演出信息"
// @Success 200 {object} resp.ResponseData "成功"
// @Router /manage/plan [post]
func (*PlanHandler)AddPlanHandler(c *gin.Context) {
	var req *dto.PlanInsertReq
	if err := c.ShouldBindJSON(&req); err!= nil {
		resp.ResponseError(c, resp.CodeInvalidParams)
		return		
	}
	// 调用service层
	if err := service.NewPlanService().AddPlan(req); err!= nil {
		resp.ResponseError(c, resp.CodeError)
		return
	}
	resp.ResponseSuccess(c, nil)
}


func (*PlanHandler)DeletePlanHandler(c *gin.Context) {
	plan_id := c.Param("plan_id")
	// 调用service层
	if err := service.NewPlanService().DeletePlan(plan_id); err!= nil {
		resp.ResponseError(c, resp.CodeError)
		return
	}
	resp.ResponseSuccess(c, nil)
}