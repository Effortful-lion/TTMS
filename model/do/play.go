package do

import "time"

type PlayStatu int

const (
	PlayStatusBefore PlayStatu = iota
	PlayStatusDuring
	PlayStatusAfter
)

type Play struct {
	PlayID          int64     `gorm:"primaryKey;autoIncrement"`
	PlayName        string    `gorm:"type:varchar(255);not null;unique"` // 剧目名称（唯一约束）
	PlayDescription string    `gorm:"type:varchar(255);not null"`        // 剧目描述
	PlayStartTime   time.Time `gorm:"type:datetime"`                     // 剧目总开始时间（如巡演周期）
	PlayEndTime     time.Time `gorm:"type:datetime"`                     // 剧目总结束时间
	PlayPrice       float64   `gorm:"type:float(10,2)"`                  // 基础票价（保留2位小数）
	PlayStatu       PlayStatu `gorm:"type:tinyint"`                      // 演出状态（枚举）

	// 关联
	Plan []Plan
}

