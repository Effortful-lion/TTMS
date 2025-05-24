package service

import (
	"TTMS/dao/mysql"
	"TTMS/dao/redis"
	"TTMS/model/dto"
	"TTMS/pkg/common"
	"errors"
)

type PlanService struct {
}

func NewPlanService() *PlanService {
	return &PlanService{}
}

func (*PlanService) GetPlan(plan_id int64) (*dto.PlanInfoResp, error) {
	plan, err := mysql.NewPlanDao().SelectPlanByID(plan_id)
	if err != nil {
		return nil, err
	}
	if plan == nil {
		return nil, errors.New("该演出计划不存在")
	}

	plan_start_str := common.ParseTimeToString( plan.PlanStartTime)
	plan_end_str := common.ParseTimeToString( plan.PlanEndTime)

	// 从 mysql 中获取剧目名称
	play, err := mysql.NewPlayDao().SelectPlayByID(plan.PlayID)
	if err!= nil {
		return nil, err
	}
	
	plan_res := &dto.PlanInfoResp{
		PlanID: plan.PlanID,
		PlayID: plan.PlayID,
		HallID: plan.HallID,
		PlayName: play.PlayName,
		PlanStartTime: plan_start_str,
		PlanEndTime: plan_end_str,
		PlanPrice: plan.PlanPrice,
		PlanStatus: int(plan.PlanStatus),
	}

	return plan_res, nil
}

func (*PlanService) GetPlanList() (*dto.PlanInfoListResp, error) {
	plan_list, err := mysql.NewPlanDao().SelectPlanList()
	if err != nil {
		return nil, err
	}
	if plan_list == nil {
		return nil, errors.New("演出计划列表为空")
	}
	
	res := &dto.PlanInfoListResp{
		PlanInfoList: make([]*dto.PlanInfoResp, 0),
	}
	for _, plan := range plan_list {
		// 从 mysql 中获取剧目名称
		play, err := mysql.NewPlayDao().SelectPlayByID(plan.PlayID)
		if err!= nil {
			return nil, err
		}
		planinfo := &dto.PlanInfoResp{
			PlanID: plan.PlanID,
			PlayID: plan.PlayID,
			HallID: plan.HallID,
			PlayName: play.PlayName,
			PlanStartTime: common.ParseTimeToString(plan.PlanStartTime),
			PlanEndTime: common.ParseTimeToString(plan.PlanEndTime),
			PlanPrice: plan.PlanPrice,
			PlanStatus: int(plan.PlanStatus),
		}
		res.PlanInfoList = append(res.PlanInfoList, planinfo)
	}
	return res, nil
}

func (*PlanService) UpdatePlan(req *dto.PlanUpdateReq) error {
	plan_id := req.PlanID
	play_id := req.PlayID
	hall_id := req.HallID
	plan_start_time := req.PlanStartTime
	plan_end_time := req.PlanEndTime
	plan_price := req.PlanPrice
	
	// 检查是否存在该计划
	plan, err := mysql.NewPlanDao().SelectPlanByID(plan_id)
	if err != nil {
		return err
	}
	if plan == nil {
		return errors.New("该演出不存在")
	}
	
	// 更新计划信息
	if err := mysql.NewPlanDao().UpdatePlan(plan_id, play_id, hall_id, plan_start_time, plan_end_time, plan_price);err != nil{
		return err
	}
	return nil
}

func (*PlanService) AddPlan(req *dto.PlanInsertReq) error {
	play_id := req.PlayID
	hall_id := req.HallID
	plan_start_time := req.PlanStartTime
	plan_end_time := req.PlanEndTime
	plan_price := req.PlanPrice

	// 插入前检查 剧目和演出厅是否存在
	play, err := mysql.NewPlayDao().SelectPlayByID(play_id)
	if err != nil {
		return err
	}
	if play == nil {
		return errors.New("剧目不存在")
	}

	hall, err := mysql.NewHallDao().SelectHall(hall_id)
	if err!= nil {
		return err
	}
	if hall == nil {
		return errors.New("演出厅不存在")
	}

	plan_id, err := mysql.NewPlanDao().InsertPlan(play_id, hall_id, plan_start_time, plan_end_time, plan_price);
	if err != nil {
		return err	
	}
	// 将 演出状态 存入 redis 中
	err = redis.RedisPlanCli.SetPlanStatusBefore(plan_id, plan_start_time, plan_end_time)
	if err != nil {
		if err.Error() == "演出开始时间已过"{
			// 数据库删除 plan
			err := mysql.NewPlanDao().DeletePlan(plan_id)
			if err != nil {
				return err
			}
			return errors.New("演出开始时间已过")
		}
		return err
	}
	return nil
}

func (*PlanService) DeletePlan(plan_id int64) error {
	err := mysql.NewPlanDao().DeletePlan(plan_id)
	if err != nil {
		return err	
	}
	err = redis.RedisPlanCli.DeletePlanStatus(plan_id)
	return err
}