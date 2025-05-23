package service

import (
	"TTMS/dao/mysql"
	"TTMS/model/dto"
	"errors"
	"fmt"
)

type HallService struct {
}

func NewHallService() *HallService {
	return &HallService{}
}

func (h *HallService) AddHall(hall_name string, hall_row, hall_col int) error {
	hall_id, err := mysql.NewHallDao().InsertHall(hall_name,hall_row,hall_col)
	if err != nil {
		return errors.New("演出厅添加失败")
	}
	return mysql.NewSeatDao().GenSeat(hall_id, hall_row, hall_col)
}

func (h *HallService) DeleteHall(hall_id int64) error {
	hall, err := mysql.NewHallDao().SelectHall(hall_id)
	if err != nil {
		return err
	}
	if hall == nil {
		return errors.New("要删除的演出厅不存在")
	} 
	err = mysql.NewHallDao().DeleteHall(hall_id)
	if err != nil {
		return err
	}
	// 删除演出厅后，对应的演出计划也要删除
	// 根据演出厅id查出对应的演出计划id列表
	ids, err := mysql.NewPlanDao().SelectPlanByHallID(hall_id)
	if err != nil {
		return err
	}
	return mysql.NewPlanDao().DeletePlanByIDs(ids)
}

func (h *HallService) UpdateHall(hall_id int64, hall_name string, hall_row, hall_col int) error {
	hall, err := mysql.NewHallDao().SelectHall(hall_id)
	if err != nil {
		return fmt.Errorf("SelectHall: %s",err.Error())
	}
	if hall == nil {
		return errors.New("要更新的演出厅不存在")
	} 
	err = mysql.NewHallDao().UpdateHall(hall_id,hall_name,hall_row, hall_col)
	if err != nil {
		return fmt.Errorf("UpdateHall: %s",err.Error())
	}
	return nil
}

func (h *HallService) GetHall(hall_id int64) (*dto.HallInfoResp, error) {
	hall, err := mysql.NewHallDao().SelectHall(hall_id)
	if err != nil {
		return nil, err
	}
	res_hall := &dto.HallInfoResp{
		HallID: hall_id,
		HallName: hall.HallName,
		HallRow: hall.HallRow,
		HallCol: hall.HallCol,
		HallTotal: hall.HallTotal,
	}
	return res_hall, nil
}

func (h *HallService) GetAllHall() (*dto.HallInfoListResp, error) {
	datas, err := mysql.NewHallDao().SelectAllHall()
	if err != nil {
		return nil, err
	}
	res := &dto.HallInfoListResp{
		Halls: make([]*dto.HallInfoResp, 0),
	}
	for _, data := range datas {
		res.Halls = append(res.Halls, &dto.HallInfoResp{
			HallID: data.HallID,
			HallName: data.HallName,
			HallRow: data.HallRow,
			HallCol: data.HallCol,
			HallTotal: data.HallTotal,
		})
	}
	return res, nil
}
