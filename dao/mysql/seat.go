package mysql

import (
	"TTMS/model/do"
	"errors"
	"fmt"
)

type SeatDao struct{}

func NewSeatDao() *SeatDao { return &SeatDao{} }

// 更新座位
func (sd *SeatDao) UpdateSeat(hall_id int64, hall_row, hall_col int) error { 
	// 查询对应 hall 的 row 和 col，如果不变就退出
	hall, err := NewHallDao().SelectHall(hall_id)
	if err != nil {return err}
	fmt.Println("原先的hall: row、col:", hall.HallRow, hall.HallCol)
	fmt.Println("更新后的hall: row、col:", hall_row, hall_col)
	if hall_row == hall.HallRow && hall_col == hall.HallCol {return nil}
	// 修改座位，对座位进行增删
	total := hall_row * hall_col
	// 增加座位
	tx := DB.Begin()
	for i := 0; i < total; i++ {
		cur_row := i / hall_col + 1
		cur_col := i % hall_col + 1
		// 如果当前座位在hall的范围内，且不存在，则创建
		if cur_row <= hall.HallRow && cur_col <= hall.HallCol {
			// 合法座位，如果不存在则创建
			seat, err := sd.SelectSeat(hall_id, cur_row, cur_col)
			if err != nil {
				tx.Rollback()
				return err
			}
			if seat == nil {
				err := sd.CreateSeat(hall_id, cur_row, cur_col)
				if err != nil {
					tx.Rollback()
					return err
				}
			}
		}
	}
	// 删除座位: 遍历所有hall_id的座位删除超出 hall 范围的座位
	err = sd.DeleteOverSeat(hall_id, hall_row, hall_col)
	if err!= nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (sd *SeatDao) DeleteOverSeat(hall_id int64, hall_row, hall_col int) error {
	// 删除座位: 遍历所有hall_id的座位删除超出 hall 范围的座位
	seats, err := sd.SelectSeatsByHallID(hall_id)
	if err!= nil {return err}
	tx := DB.Begin()
	for _, seat := range seats {
		if seat.SeatRow <= hall_row && seat.SeatCol <= hall_col {
			continue
		}
		err := tx.Delete(&do.Seat{}, seat.SeatID).Error; 
		if err!= nil {return err}
	}
	tx.Commit()
	return nil
}

func (sd *SeatDao) CreateSeat(hall_id int64, hall_row, hall_col int) error {
	// 生成单个座位
	seat := &do.Seat{HallID: hall_id, SeatRow: hall_row, SeatCol: hall_col, SeatStatus: 0}
	if err := DB.Create(seat).Error; err!= nil {
		return errors.New("座位创建失败")
	}
	return nil		
}

// 生成座位
func (sd *SeatDao) GenSeat(hall_id int64, hall_row, hall_col int) error {
	total := hall_row * hall_col
	tx := DB.Begin()
	for i := 0; i < total; i++ {
		seat := &do.Seat{HallID: hall_id, SeatRow: i / hall_col + 1, SeatCol: i % hall_col + 1, SeatStatus: 0}
		if err := tx.Create(seat).Error; err != nil {
			tx.Rollback()
			return errors.New("座位创建失败")
		}
	}
	tx.Commit()
	return nil
} 

// 选座
func (sd *SeatDao) SoldSeat(hall_id int64, row, col int) (int64, error) {
	// 选座：演出厅、行、列
	// 找到座位并判断status，如果1选座失败，如果为0改为1，选座成功
	seat_id, ok, err := sd.CheckSeat(hall_id, row, col)
	if err != nil {
		return 0, err
	}
	// 如果有座
	if ok {
		return 0, errors.New("座位已满, 选座失败")
	}
	err = DB.Model(&do.Seat{}).Where("hall_id = ? and seat_row = ? and seat_col = ?", hall_id, row, col).Update("seat_status", do.SeatstatusSold).Error
	if err != nil {
		return 0, err
	}
	return seat_id, nil
}

// 取消选座
func (sd *SeatDao) CancelSeat(hall_id int64, row, col int) error {
	_, ok, err := sd.CheckSeat(hall_id, row, col)
	if err != nil {
		return err
	}
	if !ok {
		return errors.New("座位已空, 无法取消")
	}
	return DB.Model(&do.Seat{}).Where("hall_id = ? and seat_row = ? and seat_col = ?", hall_id, row, col).Update("seat_status", do.SeatstatusNotSold).Error
}

// 查座(单个)
func (sd *SeatDao) SelectSeat(hall_id int64, row, col int)  (*do.Seat, error) { 
	seat := do.Seat{}
	err := DB.Where("hall_id = ? AND seat_row = ? AND seat_col = ?", hall_id, row, col).First(&seat).Error
	if err != nil {
		if err.Error() == "record not found" {
			return nil, nil
		}
		return nil, err
	}
	return &seat, nil
}

// 查座(多个)
func (sd *SeatDao) SelectSeatsByHallID(hall_id int64) ([]*do.Seat, error) {
	var seats []do.Seat
	err := DB.Find(&seats).Error
	if err!= nil {
		return nil, err
	}
	var res []*do.Seat
	for _, seat := range seats {
		res = append(res, &seat)
	}
	return res, nil
}

func (sd *SeatDao) SelectSeatByID(seat_id int64) (*do.Seat, error) {
	var seat do.Seat
	err := DB.Where("seat_id = ?", seat_id).First(&seat).Error
	if err != nil {
		if err.Error() == "record not found"{
			return nil, errors.New("SelectSeatByID record not found")
		}  
		return nil, err
	}
	return &seat, nil
}

// 返回 false 为没座，返回 true 为有座
func (sd *SeatDao) CheckSeat(hall_id int64, row, col int) (int64, bool, error) {
	fmt.Println("hall_id, row, col", hall_id, row, col)
	seat, err := sd.SelectSeat(hall_id, row, col)
	if err != nil { return 0, false, err}
	return seat.SeatID ,seat.SeatStatus == do.SeatstatusSold, nil
}
