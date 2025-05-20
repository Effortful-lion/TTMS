package controller

import (
	"TTMS/model/dto"
	"TTMS/pkg"
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

// @Summary 删除演出
// @Description 删除演出
// @Tags 演出管理
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param plan_id body int true "演出id"
// @Success 200 {object} resp.ResponseData "成功"
// @Router /manage/plan [delete]
func (*PlanHandler)DeletePlanHandler(c *gin.Context) {
	planStr := c.Param("plan_id")
	plan_id, err := pkg.ParseStringToInt64(planStr)
	if err!= nil {
		resp.ResponseError(c, resp.CodeInvalidParams)
		return
	}

	// 调用service层
	if err := service.NewPlanService().DeletePlan(plan_id); err!= nil {
		resp.ResponseError(c, resp.CodeError)
		return
	}
	resp.ResponseSuccess(c, nil)
}


// func (*PlanHandler)GetPlanListHandler(c *gin.Context) {
// 	// 调用service层
// 	planList, err := service.NewPlanService().GetPlanList()
// 	if err!= nil {
// 		resp.ResponseError(c, resp.CodeError)
// 		return
// 	}
// 	resp.ResponseSuccess(c, planList)
// }

