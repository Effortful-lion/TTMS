package do

type SeatStatu = int

const (
	SeatBusy SeatStatu = iota
	SeatFree
)

// Seat 座位表
type Seat struct {
	SeatID int64 `gorm:"column:seat_id;primary_key;autoIncrement"` // 座位ID，主键
	SeatRow int `gorm:"column:seat_row"` // 座位行号
	SeatCol int `gorm:"column:seat_col"` // 座位列号
	SeatStatus int `gorm:"column:seat_status"` // 座位状态，0表示空闲，1表示已售
	// 关联关系
	HallID int64 `gorm:"not null"`
	Hall   Hall  `gorm:"foreignKey:HallID"`
}