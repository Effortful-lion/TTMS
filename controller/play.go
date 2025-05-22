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

func (uc *PlayController) AddPlayHandler(c *gin.Context) {
	var req *dto.PlayInsertReq
	if err := c.ShouldBindJSON(&req); err!= nil {
		resp.ResponseError(c, resp.CodeInvalidParams)
		return
	}
	play_name := req.PlayName
	play_description := req.PlayDescription
	play_duration := req.PlayDuration
	err := service.NewPlayService().AddPlay(play_name, play_description, play_duration)
	if err != nil {
		resp.ResponseErrorWithMsg(c, resp.CodeError, err.Error())
		return
	}
	resp.ResponseSuccess(c, nil)
}

func (uc *PlayController) DeletePlayHandler(c *gin.Context) {
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

func (*PlayController) UpdatePlayHandler(c *gin.Context) {
	var req *dto.PlayUpdateReq
	if err := c.ShouldBindJSON(&req); err!= nil {
		resp.ResponseError(c, resp.CodeInvalidParams)
		return
	}
	play_id := req.PlayID
	play_name := req.PlayName
	play_description := req.PlayDescription
	play_duration := req.PlayDuration
	err := service.NewPlayService().UpdatePlay(play_id, play_name, play_description, play_duration)
	if err!= nil {
		resp.ResponseErrorWithMsg(c, resp.CodeError, err.Error())
		return
	}
	resp.ResponseSuccess(c, nil)
}

func (*PlayController) GetPlayListHandler(c *gin.Context) {
	plays, err := service.NewPlayService().GetPlayList()
	if err!= nil {
		resp.ResponseErrorWithMsg(c, resp.CodeError, err.Error())
		return
	}
	resp.ResponseSuccess(c, plays)
}

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