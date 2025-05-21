package do



type Seat struct {
	SeatID int64 `gorm:"column:seat_id;type:int;primaryKey;autoIncrement"`
	HallID int64 `gorm:"column:hall_id;type:int;not null"`
	SeatRow int `gorm:"column:seat_row;type:int;not null"`
	SeatCol int `gorm:"column:seat_col;type:int;not null"` 
	SeatStatus int8 `gorm:"column:seat_status;type:tinyint;not null"`
}

const (
	SeatstatusNotSold = iota
	SeatstatusSold
)

func (Seat) TableName() string { return "seat" }
