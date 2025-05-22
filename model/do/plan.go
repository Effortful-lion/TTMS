package do

import "time"

type PlanStatu int

const (
	PlanStatusBefore PlanStatu = iota		//  未开始
	PlanStatusDuring						//  进行中
	PlanStatusAfter							//  已结束
)

type Plan struct {
	PlanID        int64     `gorm:"primaryKey;autoIncrement"`
	PlanStartTime time.Time `gorm:"type:datetime"`    // 演出开始时间
	PlanEndTime   time.Time `gorm:"type:datetime"`    // 演出结束时间
	PlanPrice     float64   `gorm:"type:float(10,2)"` // 票价（保留2位小数）
    PlanStatus     PlanStatu `gorm:"type:tinyint"`     // 演出状态（枚举）

	PlayID int64 `gorm:"type:bigint"` // 外键，关联剧目ID
	HallID int64 `gorm:"type:bigint"` // 外键，关联演出厅ID
}

// TableName 定义表名
func (Plan) TableName() string {
	return "plan" // 对应数据库中的表名
}