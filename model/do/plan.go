package do

import "time"

type Plan struct {
	PlanID        int64 		`gorm:"primaryKey;autoIncrement"`
	PlayStartTime time.Time 	`gorm:"type:datetime"`
	PlayPrice 	  float64		`gorm:"type:float"`
	// 关联关系
	HallID int64 `gorm:"not null"`
	Hall   Hall  `gorm:"foreignKey:HallID"`

	PlayID int64 `gorm:"not null"`
	Play   Play  `gorm:"foreignKey:PlayID"`
}