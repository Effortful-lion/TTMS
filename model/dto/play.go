package dto

import "time"

type PlayStatu int

const (
	PlayStatusBefore PlayStatu = iota
	PlayStatusDuring
	PlayStatusAfter
)

type PlayInsertReq struct {
	PlayName        string    `json:"play_name" binding:"required"`
	PlayDescription string    `json:"play_description" binding:"required"`
	PlayStartTime   time.Time `json:"play_start_time" binding:"required"`
	PlayEndTime     time.Time `json:"play_end_time" binding:"required"`
	PlayPrice       float64   `json:"play_price" binding:"required"`
}

type PlayUpdateReq struct {
	PlayID          int64     `json:"play_id" binding:"required"`
	PlayName        string    `json:"play_name" binding:"required"`
	PlayDescription string    `json:"play_description" binding:"required"`
	PlayStartTime   time.Time `json:"play_start_time" binding:"required"`
	PlayEndTime     time.Time `json:"play_end_time" binding:"required"`
	PlayPrice       float64   `json:"play_price" binding:"required"`
	PlayStatu       PlayStatu `json:"play_statu" binding:"required"`
}

type PlayListResp struct {
	Plays []*PlayInfoResp `json:"plays"`
}

type PlayInfoResp struct {
	PlayID          int64     `json:"play_id"`
	PlayName        string    `json:"play_name"`
	PlayDescription string    `json:"play_description"`
	PlayStartTime   time.Time `json:"play_start_time"`
	PlayEndTime     time.Time `json:"play_end_time"`
	PlayPrice       float64   `json:"play_price"`
	PlayStatu       PlayStatu `json:"play_statu"`
}