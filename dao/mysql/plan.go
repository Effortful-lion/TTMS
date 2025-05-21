package mysql

import (
	"TTMS/model/do"
	"TTMS/pkg/common"
)

type PlanDao struct {
}

func NewPlanDao() *PlanDao {
	return &PlanDao{}
}

func (*PlanDao) SelectPlanByID(plan_id int64) (*do.Plan, error) {
	plan := do.Plan{}
	err := DB.Where("plan_id = ?", plan_id).First(&plan).Error
	if err != nil {
		if err.Error() == "record not found" {
			return nil, nil
		}
		return nil, err
	}
	return &plan, nil
}

func (*PlanDao) SelectPlanList() ([]*do.Plan, error) {
	var plans []do.Plan
	err := DB.Find(&plans).Error
	if err != nil {
		return nil, err
	}

	// 将结果转换为指针切片
	planPtrs := make([]*do.Plan, len(plans))
	for i := range plans {
		planPtrs[i] = &plans[i] // 直接取每个元素的地址，避免引用同一个变量
	}
	return planPtrs, nil
}

func (*PlanDao) UpdatePlan(plan_id, play_id, hall_id int64, plan_start_time, plan_end_time string, plan_price float64) error {
	planStartTime := common.ParseStringTime(plan_start_time)
	planEndTime := common.ParseStringTime(plan_end_time)
	plan := &do.Plan{
		PlanID: 	   plan_id,	
		PlayID:        play_id,
		HallID:        hall_id,
		PlanStartTime: planStartTime,
		PlanEndTime:   planEndTime,
		PlanPrice:     plan_price,
	}
	return DB.Save(plan).Error
}

func (*PlanDao) InsertPlan(play_id, hall_id int64, plan_start_time, plan_end_time string, plan_price float64) error {
	// 先将时间转为time.Time类型，然后再插入数据库
	planStartTime := common.ParseStringTime(plan_start_time)
	planEndTime := common.ParseStringTime(plan_end_time)
	plan := &do.Plan{
		PlayID:        play_id,
		HallID:        hall_id,
		PlanStartTime: planStartTime,
		PlanEndTime:   planEndTime,
		PlanPrice:     plan_price,
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

func (*PlanDao) DeletePlanByPlayID(play_id int64) error {
	if err := DB.Where("play_id =?", play_id).Delete(&do.Plan{}).Error; err!= nil {
		return err	
	}	
	return nil
}