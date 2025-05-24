package controller

import (
	"TTMS/model/dto"
	"TTMS/pkg/resp"
	"TTMS/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type HallController struct {

}

func NewHallController() *HallController {
	return &HallController{}
}

func (h *HallController) AddHallHandler(c *gin.Context) {
	var req *dto.HallIsertReq
	if err := c.ShouldBindJSON(&req); err != nil {
		resp.ResponseError(c, resp.CodeInvalidParams)
		return
	}
	hall_name := req.HallName
	hall_row := req.HallRow
	hall_col := req.HallCol
	err := service.NewHallService().AddHall(hall_name, hall_row, hall_col)
	if err != nil {
		resp.ResponseErrorWithMsg(c, resp.CodeError, err.Error())
		return
	}
	resp.ResponseSuccess(c, nil)
}

func (h *HallController) DeleteHallHandler(c *gin.Context) {
	hall_id := c.Param("hall_id")
	id, err := strconv.ParseInt(hall_id, 10, 64)
	if err != nil {
		resp.ResponseError(c, resp.CodeInvalidParams)
		return
	}
	err = service.NewHallService().DeleteHall(id)
	if err != nil {
		resp.ResponseError(c, resp.CodeError)
		return
	}
	resp.ResponseSuccess(c, nil)
}

func (h *HallController) UpdateHallHandler(c *gin.Context) {
	var req dto.HallUpdateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		resp.ResponseError(c, resp.CodeInvalidParams)
		return
	}
	hall_id := req.HallID
	hall_name := req.HallName
	hall_row := req.HallRow
	hall_col := req.HallCol
	err := service.NewHallService().UpdateHall(hall_id, hall_name, hall_row, hall_col)
	if err != nil {
		resp.ResponseErrorWithMsg(c, resp.CodeError, err.Error())
		return
	}
	resp.ResponseSuccess(c, nil)
}

func (h *HallController) GetHallListHandler(c *gin.Context) {
	data, err := service.NewHallService().GetAllHall()
	if err != nil {
		resp.ResponseErrorWithMsg(c, resp.CodeError, err.Error())
		return
	}
	resp.ResponseSuccess(c, data)
}

func (h *HallController) GetHallHandler(c *gin.Context) {
	hall_id := c.Param("hall_id")
	id, err := strconv.ParseInt(hall_id, 10, 64)
	if err != nil {
		resp.ResponseError(c, resp.CodeInvalidParams)
		return
	}
	data, err := service.NewHallService().GetHall(id)
	if err != nil {
		resp.ResponseErrorWithMsg(c, resp.CodeError, err.Error())
		return
	}
	resp.ResponseSuccess(c, data) 
}