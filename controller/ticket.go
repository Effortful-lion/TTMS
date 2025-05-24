package controller

import (
	"TTMS/model/dto"
	"TTMS/pkg/resp"
	"TTMS/service"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TicketController struct {
}

func NewTicketController() *TicketController {
	return &TicketController{}
}

// 单场上座率统计
func (t *TicketController) CountOnceSeatHandler(c *gin.Context) {
	strid := c.Param("plan_id")
	plan_id, err := strconv.ParseInt(strid, 10, 64)
	if err!= nil {
		resp.ResponseError(c, resp.CodeInvalidParams)
		return
	}
	data, err := service.NewTicketService().CountOnceSeat(plan_id)
	if err != nil {
		resp.ResponseErrorWithMsg(c, resp.CodeError, err.Error())
		return
	}
	resp.ResponseSuccess(c, data)
}

// 剧目上座率统计
func (t *TicketController) CountSeatHandler(c *gin.Context) {
	strid := c.Param("play_id")
	play_id, err := strconv.ParseInt(strid, 10, 64)
	if err!= nil {
		resp.ResponseError(c, resp.CodeInvalidParams)
		return
	}
	data, err := service.NewTicketService().CountSeat(play_id)
	if err != nil {
		resp.ResponseErrorWithMsg(c, resp.CodeError, err.Error())
		return
	}
	resp.ResponseSuccess(c, data)
}

// 统计票房(所有剧目)
func (t *TicketController) CountTicketListHandler(c *gin.Context) {
	// 按照 剧目 分类，统计出总共的票房
	data, err := service.NewTicketService().CountTicket()
	if err != nil {
		resp.ResponseErrorWithMsg(c, resp.CodeError, err.Error())
		return
	}
	resp.ResponseSuccess(c, data)
}

func (t *TicketController) CountOnceTicketPercentageHandler(c *gin.Context) {
	strid := c.Param("plan_id")
	plan_id, err := strconv.ParseInt(strid, 10, 64)
	if err!= nil {
		resp.ResponseError(c, resp.CodeInvalidParams)
		return
	}
	data, err := service.NewTicketService().CountOnceTicketPercentageByID(plan_id)
	if err != nil {
		resp.ResponseErrorWithMsg(c, resp.CodeError, err.Error())
		return
	}
	resp.ResponseSuccess(c, data)
}

// 统计票房（单个剧目的单场演出）
func (t *TicketController) CountOnceTicketHandler(c *gin.Context) {
	strid := c.Param("plan_id")
	plan_id, err := strconv.ParseInt(strid, 10, 64)
	if err!= nil {
		resp.ResponseError(c, resp.CodeInvalidParams)
		return
	}
	data, err := service.NewTicketService().CountOnceTicketByID(plan_id)
	if err != nil {
		resp.ResponseErrorWithMsg(c, resp.CodeError, err.Error())
		return
	}
	resp.ResponseSuccess(c, data)
}


// 统计票房（单个剧目）
func (t *TicketController) CountTicketHandler(c *gin.Context) {
	strid := c.Param("play_id")
	play_id, err := strconv.ParseInt(strid, 10, 64)
	if err!= nil {
		resp.ResponseError(c, resp.CodeInvalidParams)
		return
	}
	data, err := service.NewTicketService().CountTicketByID(play_id)
	if err != nil {
		resp.ResponseErrorWithMsg(c, resp.CodeError, err.Error())
		return
	}
	resp.ResponseSuccess(c, data)
}

// 核销票
func (t *TicketController) VerifyHandler(c *gin.Context) {
	var req *dto.TicketVerifyReq
	if err := c.ShouldBindJSON(&req); err != nil {
		resp.ResponseError(c, resp.CodeInvalidParams)
		return
	}
	err := service.NewTicketService().VerifyTicket(req.TicketID)
	if err != nil {
		resp.ResponseErrorWithMsg(c, resp.CodeError, err.Error())
		return
	}
	resp.ResponseSuccess(c, nil)
}

// 查票
func (t *TicketController) GetTicketListHandler(c *gin.Context) {
	customer_id := GetCurrentUserID(c)
	list, err := service.NewTicketService().GetTicketList(customer_id)
	if err != nil {
		resp.ResponseErrorWithMsg(c, resp.CodeError, err.Error())
		return
	}
	resp.ResponseSuccess(c, list)
}

// 退票
func (t *TicketController) CancelHandler(c *gin.Context) {
	var req dto.TicketCancelReq
	if err := c.ShouldBindJSON(&req); err != nil { 
		resp.ResponseError(c, resp.CodeInvalidParams)
		return
	}
	fmt.Println(req.TicketID)
	err := service.NewTicketService().CancelTicket(&req)
	if err != nil {
		resp.ResponseErrorWithMsg(c, resp.CodeError, err.Error())
		return
	}
	resp.ResponseSuccess(c, nil)
}

// 买票
func (t *TicketController) BuyHandler(c *gin.Context) {
	var req *dto.TicketBuyReq
	if err := c.ShouldBindJSON(&req); err != nil {
		resp.ResponseError(c, resp.CodeInvalidParams)
		return
	}
	customerID := GetCurrentUserID(c)
	auth := GetCurrentUserAuthority(c)
	err := service.NewTicketService().BuyTicket(customerID, auth, req)
	if err != nil {
		resp.ResponseErrorWithMsg(c, resp.CodeError, err.Error())
		return
	}
	resp.ResponseSuccess(c, nil)
}