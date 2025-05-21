package service

import (
	"TTMS/dao/mysql"
	"TTMS/model/dto"
	"errors"
)

type HallService struct {
}

func NewHallService() *HallService {
	return &HallService{}
}

func (h *HallService) AddHall(hall_name string, hall_row, hall_col int) error {
	return mysql.NewHallDao().InsertHall(hall_name,hall_row,hall_col)
}

func (h *HallService) DeleteHall(hall_id int64) error {
	hall, err := mysql.NewHallDao().SelectHall(hall_id)
	if err != nil {
		return err
	}
	if hall == nil {
		return errors.New("要删除的演出厅不存在")
	} 
	return mysql.NewHallDao().DeleteHall(hall_id)
}

func (h *HallService) UpdateHall(hall_id int64, hall_name string, hall_row, hall_col int) error {
	hall, err := mysql.NewHallDao().SelectHall(hall_id)
	if err != nil {
		return err
	}
	if hall == nil {
		return errors.New("要更新的演出厅不存在")
	} 
	return mysql.NewHallDao().UpdateHall(hall_id,hall_name,hall_row, hall_col)
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
