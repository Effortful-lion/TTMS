package dto



type PlanInsertReq struct {
	PlayID int64 `json:"play_id" binding:"required"` // 剧目ID，唯一
	HallID int64 `json:"hall_id" binding:"required"` // 演出厅ID，唯一
	PlanStartTime string `json:"plan_start_time" binding:"required"` // 演出计划开始时间，唯一
	PlanEndTime string `json:"plan_end_time" binding:"required"` // 演出计划结束时间，唯一
	PlanPrice float64 `json:"plan_price" binding:"required"` // 演出计划价格，唯一
}

type PlanUpdateReq struct {
	PlanID int64 `json:"plan_id" binding:"required"` // 计划ID，唯一
	PlayID int64 `json:"play_id" binding:"required"` // 剧目ID，唯一
	HallID int64 `json:"hall_id" binding:"required"` // 演出厅ID，唯一
	PlanStartTime string `json:"plan_start_time" binding:"required"` // 演出计划开始时间，唯一
	PlanEndTime string `json:"plan_end_time" binding:"required"` // 演出计划结束时间，唯一
	PlanPrice float64 `json:"plan_price" binding:"required"` // 演出计划价格，唯一
}

type PlanInfoResp struct {
	PlanID int64 `json:"plan_id"` // 计划ID，唯一
	PlayID int64 `json:"play_id"` // 剧目ID，唯一
	HallID int64 `json:"hall_id"` // 演出厅ID，唯一
	PlayName  string `json:"play_name"`	// 剧目名称，唯一
	PlanStartTime string `json:"plan_start_time"` // 演出计划开始时间，唯一
	PlanEndTime string `json:"plan_end_time"` // 演出计划结束时间，唯一
	PlanPrice float64 `json:"plan_price"` // 演出计划价格，唯一
	PlanStatus int `json:"plan_status"` // 演出计划状态，唯一
}

type PlanInfoListResp struct {
	PlanInfoList []*PlanInfoResp `json:"plan_info_list"`
}