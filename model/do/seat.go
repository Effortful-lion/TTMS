package do

type SeatStatu = int

const (
	SeatBusy SeatStatu = iota
	SeatFree
)


type Seat struct {
	SeatID     int64 `gorm:"column:seat_id;primaryKey;autoIncrement"` // 座位ID，主键
	SeatRow    int   `gorm:"column:seat_row;not null"`                // 座位行号
	SeatCol    int   `gorm:"column:seat_col;not null"`                // 座位列号
	SeatStatus int   `gorm:"column:seat_status;not null;default:0"`   // 座位状态（0:空闲，1:已售）

	// 依赖
	Hall   Hall
	HallID int64 `gorm:"foreignKey:HallID;references:HallID;index;not null"` // 演出厅ID，外键
}