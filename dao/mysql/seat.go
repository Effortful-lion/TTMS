package mysql

import (
	"TTMS/model/do"
	"errors"
)

type SeatDAO struct {
}

func NewSeatDao() *SeatDAO {
	return &SeatDAO{}
}

// 控制座位的状态为 “有座”
func (*SeatDAO) ToBusy(hall_id int64, row, col int) error {
	seat := &do.Seat{
		SeatRow:    row,
		SeatCol:    col,
		SeatStatus: do.SeatBusy,
		HallID:     hall_id,
	}
	err := DB.Save(seat).Error
	if err != nil {
		return errors.New("座位状态更新失败")
	}
	return nil
}

func (*SeatDAO) ToFree(hall_id int64, row, col int) error {
	seat := &do.Seat{
		SeatRow:    row,
		SeatCol:    col,
		SeatStatus: do.SeatFree,
		HallID:     hall_id,
	}
	err := DB.Save(seat).Error
	if err != nil {
		return errors.New("座位状态更新失败")
	}
	return nil
}

// （管理员）批量生成座位
func (*SeatDAO) BatchInsertSeat(hall_id int64, row, col int) error {
	total := row * col
	var seats []*do.Seat
	for i := 0; i < total; i++ {
		// 这里需要计算正确的行和列，当前代码逻辑有误，下面是修正示例
		seatRow := i/col + 1
		seatCol := i%col + 1
		seat := &do.Seat{
			SeatRow:    seatRow,
			SeatCol:    seatCol,
			HallID:     hall_id,
			SeatStatus: do.SeatFree, // 假设初始状态为空闲
		}
		seats = append(seats, seat)
	}
	// 批量插入座位数据
	err := DB.CreateInBatches(seats, total).Error
	if err != nil {
		return errors.New("批量插入座位失败")
	}
	return nil
}
