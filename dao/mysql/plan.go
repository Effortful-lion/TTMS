package mysql

import (
	"TTMS/model/do"
	"TTMS/pkg"
)

type PlanDao struct {
}

func NewPlanDao() *PlanDao {
	return &PlanDao{}
}

func (*PlanDao) InsertPlan(play_id, hall_id int64, plan_start_time, plan_end_time string, plan_price float64, plan_status int) error {
	// 先将时间转为time.Time类型，然后再插入数据库
	planStartTime := pkg.ParseStringTime(plan_start_time)
	planEndTime := pkg.ParseStringTime(plan_end_time)
	plan := &do.Plan{
		PlayID:          play_id,
		HallID:          hall_id,
		PlanStartTime:   planStartTime,
		PlanEndTime:     planEndTime,
		PlanPrice:       plan_price,
		PlanStatu:      do.PlanStatu(plan_status),
	}
	if err := DB.Create(&plan).Error; err != nil {
		return err
	}
	return nil
}

func (*PlanDao) DeletePlan(plan_id int64) error {
	if err := DB.Delete(&do.Plan{}, plan_id).Error; err!= nil {
		return err
	}else{
		return nil
	}	
}