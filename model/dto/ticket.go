package dto

import "time"

type TicketBuyReq struct {
	PlanID  int64 `json:"plan_id" binding:"required"`
	SeatRow int   `json:"seat_row" binding:"required"`
	SeatCol int   `json:"seat_col" binding:"required"`
}

type TicketCancelReq struct {
	TicketID int64 `json:"ticket_id" binding:"required"`
}

type TicketInfoResp struct {
	// 票的相关信息
	TicketID         int64     `json:"ticket_id"`
	CustomerID       int64     `json:"customer_id"`
	SeatID           int64     `json:"seat_id"`
	PlanID           int64     `json:"plan_id"`
	PlayID 			 int64     `json:"play_id"`
	// 票面展示信息
	CustomerName     string    `json:"customer_name"`
	SeatRow          int       `json:"seat_row"`
	SeatCol          int       `json:"seat_col"`
	PlayName         string    `json:"play_name"`
	HallName         string    `json:"hall_name"`
	TicketPrice      float64   `json:"ticket_price"`
	TicketStatus     int8       `json:"ticket_status"`
	TicketExpireTime time.Time `json:"ticket_expire_time"`
}

type TicketInfoListResp struct {
	// 票列表
	Tickets []*TicketInfoResp `json:"tickets"`
}

type TicketVerifyReq struct {
	TicketID int64 `json:"ticket_id" binding:"required"`
}

// 统计票房，分类返回票房
type TicketCountResp struct {
	PlayName    string  `json:"play_name"`		// 剧目名
	TotalMoney  float64 `json:"total_money"`		// 票房
	CountTime   time.Time   `json:"count_time"`			// 统计时间
}

type TicketCountListResp struct { 
	TicketCountList []TicketCountResp `json:"ticket_count_list"`
}

type TicketPlanPercentageResp struct {
	PlanID        int64   `json:"plan_id"`		// 场次ID
	Percentage    float64 `json:"percentage"`		// 占比
}

type TicketPayReq struct {
	CustomerID int64 `json:"customer_id"`
	PlanID int64 `json:"plan_id"`
	SeatID int64 `json:"seat_id"`
	Money float64 `json:"money"`
}

type TicketPayResp struct {
	CustomerID int64 `json:"customer_id"`
	PlanID int64 `json:"plan_id"`
	SeatID int64 `json:"seat_id"`
	Money float64 `json:"money"`
}