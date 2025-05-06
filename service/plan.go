package service

import (
	"TTMS/dao/mysql"
	"TTMS/dao/redis"
	"TTMS/model/dto"
	"TTMS/pkg"
)

type PlanService struct {
}

func NewPlanService() *PlanService {
	return &PlanService{}
}

func (*PlanService) AddPlan(req *dto.PlanInsertReq) error {
	play_id := req.PlayID
	hall_id := req.HallID
	plan_start_time := req.PlanStartTime
	plan_end_time := req.PlanEndTime
	plan_price := req.PlanPrice
	plan_status := req.PlanStatus

	if err := mysql.NewPlanDao().InsertPlan(play_id, hall_id, plan_start_time, plan_end_time, plan_price, plan_status);err != nil {
		return err	
	}
	// 将 演出状态 存入 redis 中
	st := pkg.ParseStringTimeToTimeStamp(plan_start_time)
	et := pkg.ParseStringTimeToTimeStamp(plan_end_time)
	expire := et - st
	if err := redis.NewPlanRedis().SetPlanStatus(play_id, plan_status,expire);err!= nil {
		return err
	}
	return nil 
}