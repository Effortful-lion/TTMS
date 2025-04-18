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
	PyayStatu 		PlayStatu			`gorm:"type:tinyint"`
}

