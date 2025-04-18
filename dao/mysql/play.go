package mysql

import (
	"TTMS/model/do"
	"time"
)

// 剧目的dao层
type PlayDao struct{}

func NewPlayDao() *PlayDao {
	return &PlayDao{}
}

func (*PlayDao)InsertPlay(play_name, play_description string, play_start, play_end time.Time, play_price float64) error {
	play := &do.Play{
		PlayName: play_name,
		PlayDescription: play_description,
		PlayStartTime: play_start,
		PlayEndTime: play_end,
		PlayPrice: play_price,
	}
	return DB.Create(play).Error
}

func (*PlayDao)SelectAllPlay() ([]*do.Play, error) {
	var plays []*do.Play
	err := DB.Find(&plays).Error
	return plays, err
}

func (*PlayDao)SelectPlayByID(play_id int64) (*do.Play, error) {
	var play *do.Play
	err := DB.Where("play_id = ?", play_id).First(&play).Error	
	return play, err
}

func (*PlayDao)UpdatePlay(play_id int,play_name, play_description string, play_start, play_end time.Time, play_price float64, play_status int) error {
	play := &do.Play{
		PlayID: int64(play_id),
		PlayName: play_name,
		PlayDescription: play_description,
		PlayStartTime: play_start,
		PlayEndTime: play_end,
		PlayPrice: play_price,
		PyayStatu: do.PlayStatu(play_status),	
	}
	return DB.Save(play).Error
}

func (*PlayDao)DeletePlay(play_id int64) error {
	var play *do.Play
	return DB.Where("play_id =?", play_id).Delete(&play).Error
}