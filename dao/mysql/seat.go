package mysql

import (
	"TTMS/model/do"
	"errors"
)

type SeatDao struct{}

func NewSeatDao() *SeatDao { return &SeatDao{} }

// 选座
func (sd *SeatDao) SoldSeat(hall_id int64, row, col int) error {
	// 选座：演出厅、行、列
	// 找到座位并判断status，如果1选座失败，如果为0改为1，选座成功
	ok, err := sd.CheckSeat(hall_id, row, col)
	if err != nil {
		return err
	}
	// 如果有座
	if ok {
		return errors.New("座位已满, 选座失败")
	}
	return DB.Model(&do.Seat{}).Where("hall_id = ? and seat_row = ? and seat_col = ?", hall_id, row, col).Update("seat_status", do.SeatstatusSold).Error
}

// 取消选座
func (sd *SeatDao) CancelSeat(hall_id int64, row, col int) error {
	ok, err := sd.CheckSeat(hall_id, row, col)
	if err != nil {
		return err
	}
	if !ok {
		return errors.New("座位已空, 无法取消")
	}
	return DB.Model(&do.Seat{}).Where("hall_id = ? and seat_row = ? and seat_col = ?", hall_id, row, col).Update("seat_status", do.SeatstatusNotSold).Error
}

func (sd *SeatDao) SelectSeat(hall_id int64, row, col int)  (*do.Seat, error) { 
	seat := do.Seat{}
	err := DB.Where("hall_id = ? AND seat_row = ? AND seat_col = ?", hall_id, row, col).First(&seat).Error
	if err != nil { return nil, err}
	return &seat, nil
}

// 返回 false 为没座，返回 true 为有座
func (sd *SeatDao) CheckSeat(hall_id int64, row, col int) (bool, error) {
	seat, err := sd.SelectSeat(hall_id, row, col)
	if err != nil { return false, err}
	return seat.SeatStatus == do.SeatstatusSold, nil
}
