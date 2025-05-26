package controller

import (
	"TTMS/model/dto"
	"TTMS/pkg/common"
	"TTMS/pkg/resp"
	"TTMS/service"
	"fmt"

	"github.com/gin-gonic/gin"
)

type PlanController struct {
}

func NewPlanController() *PlanController {
	return &PlanController{}
}

func (uc *PlanController) GetPlanHandler(c *gin.Context) {
	planStr := c.Param("plan_id")
	plan_id, err := common.ParseStringToInt64(planStr)
	if err!= nil {
		resp.ResponseError(c, resp.CodeInvalidParams)
		return
	}
	// 调用service层
	plan, err := service.NewPlanService().GetPlan(plan_id)
	if err!= nil {
		resp.ResponseErrorWithMsg(c, resp.CodeError, err.Error())
		return
	}
	resp.ResponseSuccess(c, plan)
}

func (uc *PlanController) GetPlanListHandler(c *gin.Context) {
	// 调用service层
	planList, err := service.NewPlanService().GetPlanList()
	if err!= nil {
		resp.ResponseError(c, resp.CodeError)
		return
	}
	resp.ResponseSuccess(c, planList)
}

func (uc *PlanController) AddPlanHandler(c *gin.Context) {
	var req dto.PlanInsertReq
	if err := c.ShouldBindJSON(&req); err!= nil {
		resp.ResponseError(c, resp.CodeInvalidParams)
		return		
	}
	fmt.Println("play_id: ", req.PlayID, "plan_start_time: ", req.PlanStartTime, "plan_end_time: ", req.PlanEndTime, "plan_price: ", req.PlanPrice, "hall_id: ", req.HallID)
	// 调用service层
	if err := service.NewPlanService().AddPlan(&req); err!= nil {
		resp.ResponseErrorWithMsg(c, resp.CodeError, err.Error())
		return
	}
	resp.ResponseSuccess(c, nil)
}

func (uc *PlanController)DeletePlanHandler(c *gin.Context) {
	planStr := c.Param("plan_id")
	plan_id, err := common.ParseStringToInt64(planStr)
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

func (uc *PlanController)UpdatePlanHandler(c *gin.Context) {
	var req dto.PlanUpdateReq
	if err := c.ShouldBindJSON(&req); err!= nil {
		resp.ResponseError(c, resp.CodeInvalidParams)
		return
	}
	// 调用service层
	if err := service.NewPlanService().UpdatePlan(&req); err!= nil {
		resp.ResponseErrorWithMsg(c, resp.CodeError, err.Error())
		return
	}
	resp.ResponseSuccess(c, nil)
}