package controller

import (
	"TTMS/model/dto"
	"TTMS/pkg/resp"
	"TTMS/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PlayController struct {
}

func NewPlayController() *PlayController {
	return &PlayController{}
}

// @Summary 添加剧目
// @Description 添加剧目
// @Tags 剧目管理
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param object body dto.PlayInsertReq true "剧目信息"
// @Success 200 {object} resp.ResponseData "成功"
// @Router /manage/play [post]
func (*PlayController) AddPlayHandler(c *gin.Context) {
	var req *dto.PlayInsertReq
	if err := c.ShouldBindJSON(&req); err!= nil {
		resp.ResponseError(c, resp.CodeInvalidParams)
		return
	}
	play_name := req.PlayName
	play_description := req.PlayDescription
	play_start_time := req.PlayStartTime
	play_end_time := req.PlayEndTime
	play_price := req.PlayPrice
	err := service.NewPlayService().AddPlay(play_name, play_description, play_start_time, play_end_time, play_price)
	if err != nil {
		resp.ResponseErrorWithMsg(c, resp.CodeError, err.Error())
		return
	}
	resp.ResponseSuccess(c, nil)
}

// @Summary 删除剧目
// @Description 删除剧目
// @Tags 剧目管理
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param play_id body int true "剧目id"
// @Success 200 {object} resp.ResponseData "成功"
// @Router /manage/play/{play_id} [delete]
func (*PlayController) DeletePlayHandler(c *gin.Context) {
	strid := c.Param("play_id")
	play_id, err := strconv.ParseInt(strid, 10, 64)
	if err!= nil {
		resp.ResponseError(c, resp.CodeInvalidParams)
		return
	}
	err = service.NewPlayService().DeletePlay(play_id)
	if err!= nil {
		resp.ResponseErrorWithMsg(c, resp.CodeError, err.Error())
		return
	}
	resp.ResponseSuccess(c, nil)
}

// @Summary 更新剧目
// @Description 更新剧目
// @Tags 剧目管理
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param object body dto.PlayUpdateReq true "剧目更新信息"
// @Success 200 {object} resp.ResponseData "成功"
// @Router /manage/play [put]
func (*PlayController) UpdatePlayHandler(c *gin.Context) {
	var req *dto.PlayUpdateReq
	if err := c.ShouldBindJSON(&req); err!= nil {
		resp.ResponseError(c, resp.CodeInvalidParams)
		return
	}
	play_id := req.PlayID
	play_name := req.PlayName
	play_description := req.PlayDescription
	play_start_time := req.PlayStartTime
	play_end_time := req.PlayEndTime
	play_price := req.PlayPrice
	err := service.NewPlayService().UpdatePlay(play_id, play_name, play_description, play_start_time, play_end_time, play_price)
	if err!= nil {
		resp.ResponseErrorWithMsg(c, resp.CodeError, err.Error())
		return
	}
	resp.ResponseSuccess(c, nil)
}

// @Summary 获得所有剧目
// @Description 获得所有剧目
// @Tags 剧目管理
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Success 200 {object} resp.ResponseData "成功"
// @Router /manage/play [get]
func (*PlayController) GetPlayListHandler(c *gin.Context) {
	plays, err := service.NewPlayService().GetPlayList()
	if err!= nil {
		resp.ResponseErrorWithMsg(c, resp.CodeError, err.Error())
		return
	}
	resp.ResponseSuccess(c, plays)
}

// @Summary 获得特定剧目
// @Description 获得特定剧目
// @Tags 剧目管理
// @Accept application/json
// @Produce application/json
// @Param Authorization header string true "Bearer 用户令牌"
// @Param play_id body int true "剧目id"
// @Success 200 {object} resp.ResponseData "成功"
// @Router /manage/play/{play_id} [get]
func (*PlayController) GetPlayHandler(c *gin.Context) {
	strid := c.Param("play_id")
	play_id, err := strconv.ParseInt(strid, 10, 64)
	if err!= nil {
		resp.ResponseError(c, resp.CodeInvalidParams)
		return
	}
	play, err := service.NewPlayService().GetPlay(play_id)
	if err!= nil {
		resp.ResponseErrorWithMsg(c, resp.CodeError, err.Error())
		return	
	}
	resp.ResponseSuccess(c, play)
}