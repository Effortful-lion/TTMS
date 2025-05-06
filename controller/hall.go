package controller

import (
	"TTMS/model/dto"
	"TTMS/pkg/resp"
	"TTMS/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type HallHandler struct {
}

func NewHallController() *HallHandler {
	return &HallHandler{}
}

// @Summary 添加演出厅
// @Description 添加演出厅
// @Tags 演出厅管理
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param object body dto.HallIsertReq true "演出厅信息"
// @Success 200 {object} resp.ResponseData "成功"
// @Router /manage/hall [post]
func (h *HallHandler) AddHallHandler(c *gin.Context) {
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
	}
	resp.ResponseSuccess(c, nil)
}

// @Summary 删除演出厅
// @Description 删除演出厅
// @Tags 演出厅管理
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param hall_id body int true "演出厅id"
// @Success 200 {object} resp.ResponseData "成功"
// @Router /manage/hall [delete]
func (h *HallHandler) DeleteHallHandler(c *gin.Context) {
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

// @Summary 更新演出厅
// @Description 更新演出厅
// @Tags 演出厅管理
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param object body dto.HallUpdateReq true "演出厅信息"
// @Success 200 {object} resp.ResponseData "成功"
// @Router /manage/hall [put]
func (h *HallHandler) UpdateHallHandler(c *gin.Context) {
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

// @Summary 查询所有演出厅
// @Description 查询所有演出厅
// @Tags 演出厅管理
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Success 200 {object} resp.ResponseData "成功"
// @Router /manage/hall [get]
func (h *HallHandler) GetHallListHandler(c *gin.Context) {
	data, err := service.NewHallService().GetAllHall()
	if err != nil {
		resp.ResponseErrorWithMsg(c, resp.CodeError, err.Error())
		return
	}
	resp.ResponseSuccess(c, data)
}

// @Summary 查询所有演出厅
// @Description 查询所有演出厅
// @Tags 演出厅管理
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param hall_id body int true "演出厅id"
// @Success 200 {object} resp.ResponseData "成功"
// @Router /manage/hall/{hall_id} [get]
func (h *HallHandler) GetHallHandler(c *gin.Context) {
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

