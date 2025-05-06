package do


type Ticket struct {
	TicketID     int64   `gorm:"primaryKey;autoIncrement"`
	TicketPrice  float64 `gorm:"type:float(10,2);not null"`                  // 实际支付票价（可能有折扣）
	TicketStatus string  `gorm:"type:varchar(20);not null;default:'unpaid'"` // 票状态（unpaid/paid/refunded）

	// 依赖
	Plan       Plan
	PlanID     int64 `gorm:"foreignKey:PlanID;references:PlanID;index"` // 演出计划ID，外键
	Seat       Seat
	SeatID     int64 `gorm:"foreignKey:SeatID;references:SeatID;index"` // 座位ID，外键
	UserInfo   UserInfo
	UserInfoID int64 `gorm:"foreignKey:UserInfoID;references:UserInfoID;index;not null"` // 用户ID，外键
}