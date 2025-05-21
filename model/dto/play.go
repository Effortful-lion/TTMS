package dto



type PlayInsertReq struct {
	PlayName        string    `json:"play_name" binding:"required"`
	PlayDescription string    `json:"play_description" binding:"required"`
	PlayDuration    int       `json:"play_duration" binding:"required"`
	PlayPrice       float64   `json:"play_price" binding:"required"`
}

type PlayUpdateReq struct {
	PlayID          int64     `json:"play_id" binding:"required"`
	PlayName        string    `json:"play_name" binding:"required"`
	PlayDescription string    `json:"play_description" binding:"required"`
	PlayDuration    int       `json:"play_duration" binding:"required"`
	PlayPrice       float64   `json:"play_price" binding:"required"`
}

type PlayInfoResp struct {
	PlayID          int64     `json:"play_id"`
	PlayName        string    `json:"play_name"`
	PlayDescription string    `json:"play_description"`
	PlayDuration    int       `json:"play_duration"`
	PlayPrice       float64   `json:"play_price"`
}

type PlayListResp struct {
	Plays []*PlayInfoResp `json:"plays"`
}