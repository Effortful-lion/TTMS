package do

import "time"

type PlayStatu int

const (
	PlayStatusBefore PlayStatu = iota
	PlayStatusDuring
	PlayStatusAfter
)

type Play struct {
	PlayID          int64				`gorm:"primaryKey;autoIncrement"`
	PlayName        string				`gorm:"type:varchar(255);not null"`
	PlayDescription string				`gorm:"type:varchar(255);not null"`
	PlayStartTime 	time.Time			`gorm:"type:datetime"`
	PlayEndTime		time.Time			`gorm:"type:datetime"`
	PlayPrice 		float64				`gorm:"type:float"`
	PlayStatu 		PlayStatu			`gorm:"type:tinyint"`

	// 依赖关系：一个剧目对应多个演出计划，一个演出计划对应多个剧目（一对多）
	Plans            []Plan				 `gorm:"foreignKey:PlayID"`
}

