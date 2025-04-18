package do


type Ticket struct {
	// 基本信息
	TicketID int64 `gorm:"primaryKey;autoIncrement"`
	TicketPrice float64 `gorm:"not null"`
	TicketStatus string `gorm:"not null"`
	
	// 关联信息
	UserID int64 `gorm:"not null"`				// 一票一客户，一个客户可以买多张票（一对多）
	User UserInfo `gorm:"foreignKey:UserID"`	// 关联到 UserInfo 表的 UserID 字段

	PlayID int64 `gorm:"not null"`				// 一票一演出，一个演出可以有多张票（一对多）
	Play Play `gorm:"foreignKey:PlayID"`		// 关联到 Play 表的 PlayID 字段

	SeatID int64 `gorm:"not null"`				// 一票一座位，一个座位只能被一张票占据（一对一）
	Seat Seat `gorm:"foreignKey:SeatID"`		// 关联到 Seat 表的 SeatID 字段
}