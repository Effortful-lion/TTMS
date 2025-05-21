package do

import "time"

const (
	// 已支付+未过期 票的使用状态
	TicketStatusUnUsed = iota 	// 未核销
	TicketStatusUsed			// 已核销
	TicketStatusCanceled		// 已取消（未核销+票过期、已核销）
)

type Ticket struct {
	TicketID     int64    				 `gorm:"column:ticket_id;type:int;primaryKey;autoIncrement"`
	CustomerID   int64    				`gorm:"column:customer_id;type:int;not null"`
	CustomerName string 				`gorm:"column:customer_name;type:varchar(100);not null"`
	SeatID       int64    				 `gorm:"column:seat_id;type:int;not null"`
	PlanID       int64     				`gorm:"column:plan_id;type:int;not null"`
	TicketPrice  float64  				 `gorm:"column:ticket_price;type:float(10,2);not null"`
	TicketStatus int8     				 `gorm:"column:ticket_status;type:tinyint;not null"`
	TicketExpireTime   time.Time 		`gorm:"column:ticket_expire_time;type:datetime;not null"`
}

func (Ticket) TableName() string {
	return "ticket"
}
