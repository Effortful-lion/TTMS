package do

import "time"

type PlanStatu int

const (
	PlanStatusBefore PlanStatu = iota
	PlanStatusDuring
	PlanStatusAfter
)

type Plan struct {
	PlanID        int64     `gorm:"primaryKey;autoIncrement"`
	PlanStartTime time.Time `gorm:"type:datetime"`    // 演出开始时间
	PlanEndTime   time.Time `gorm:"type:datetime"`    // 演出结束时间
	PlanPrice     float64   `gorm:"type:float(10,2)"` // 票价（保留2位小数）
    PlanStatu     PlanStatu `gorm:"type:tinyint"`     // 演出状态（枚举）

	// 依赖
	Play   Play
	PlayID int64 `gorm:"foreignKey:PlayID;references:PlayID;index;not null"` // 演出ID，外键
	Hall   Hall
	HallID int64 `gorm:"foreignKey:HallID;references:HallID;index;not null"` // 演出厅ID，外键

	// 关联
	Tickets []Ticket
}