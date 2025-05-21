package mysql

import (
	"TTMS/model/do"
	"errors"
)

type HallDao struct{}

func NewHallDao() *HallDao {
	return &HallDao{}
}

func (*HallDao) InsertHall(hall_name string, hall_row, hall_col int) (err error) {
	hall := &do.Hall{
		HallName: hall_name,
		HallRow:  hall_row,
		HallCol:  hall_col,
		HallTotal: hall_row * hall_col,
	}
	if err = DB.Create(&hall).Error; err != nil {
		return errors.New("创建失败")	
	}
	return nil
}

func (*HallDao) DeleteHall(hall_id int64) (err error) {
	if err = DB.Delete(&do.Hall{}, hall_id).Error; err!= nil {
		return errors.New("删除失败")
	}
	return nil
}

func (*HallDao) UpdateHall(hall_id int64, hall_name string, hall_row, hall_col int) error {
	hall := &do.Hall{
		HallID: int64(hall_id),
		HallName: hall_name,
		HallRow:  hall_row,
		HallCol:  hall_col,
		HallTotal: hall_row * hall_col,
	}
	if err := DB.Save(hall).Error; err != nil {
		return errors.New("更新失败")
	}
	return nil
}

func (*HallDao) SelectHall(hall_id int64) (hall *do.Hall, err error) {
	var h do.Hall
	if err := DB.Where("hall_id = ?", hall_id).First(&h).Error; err != nil {
		if err.Error() == "record not found" {
			return nil, nil	
		}
		return nil, err
	}
	return &h, nil
}

func (*HallDao) SelectAllHall() ([]*do.Hall, error) {
	var hs []do.Hall
	err := DB.Find(&hs).Error
	if err != nil {
		return nil, err
	}
	// 将非指针转换为指针类型
	hsptr := make([]*do.Hall, len(hs))
	for i := range hs {
		hsptr[i] = &hs[i]
	}
	return hsptr, nil
}
