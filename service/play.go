package service

import (
	"TTMS/dao/mysql"
	"TTMS/model/dto"
	"time"

	"gorm.io/gorm"
)

type PlayService struct {
}

func NewPlayService() *PlayService {
	return &PlayService{}
}

func (*PlayService) GetPlay(play_id int64) (data *dto.PlayInfoResp, err error) {
	dd := mysql.NewPlayDao()
	play, err := dd.SelectPlayByID(play_id)
	if err != nil {
		return nil, err
	}
	return &dto.PlayInfoResp{
		PlayID:          play.PlayID,
		PlayName:        play.PlayName,
		PlayDescription: play.PlayDescription,
		PlayStartTime:   play.PlayStartTime,
		PlayEndTime:     play.PlayEndTime,
		PlayPrice:       play.PlayPrice,
	}, nil
}

func (*PlayService) GetPlayList() (data *dto.PlayListResp, err error) {
	dd := mysql.NewPlayDao()
	plays, err := dd.SelectAllPlay()
	if err != nil {
		return nil, err
	}
	// 初始化 playList 指针
	playList := &dto.PlayListResp{
		Plays: make([]*dto.PlayInfoResp, 0),
	}
	for _, play := range plays {
		play_info := &dto.PlayInfoResp{
			PlayID:          play.PlayID,
			PlayName:        play.PlayName,
			PlayDescription: play.PlayDescription,
			PlayStartTime:   play.PlayStartTime,
			PlayEndTime:     play.PlayEndTime,
			PlayPrice:       play.PlayPrice,
		}
		playList.Plays = append(playList.Plays, play_info)
	}
	return playList, nil
}

func (*PlayService) AddPlay(play_name, play_description string, play_start, play_end time.Time, play_price float64) (err error) {
	return mysql.NewPlayDao().InsertPlay(play_name, play_description, play_start, play_end, play_price)
}

func (*PlayService) DeletePlay(play_id int64) (err error) {
	// 检查 play_id 是否存在于数据库中
	play, err := mysql.NewPlayDao().SelectPlayByID(play_id)
	if err!= nil {
		return err
	}
	if play == nil {
		return gorm.ErrRecordNotFound	
	}
	return mysql.NewPlayDao().DeletePlay(play_id)
}

func (*PlayService) UpdatePlay(play_id int64, play_name, play_description string, play_start, play_end time.Time, play_price float64) (err error) {
	// 检查 play_id 是否存在于数据库中
	play, err := mysql.NewPlayDao().SelectPlayByID(play_id)
	if err!= nil {
		return err
	}
	if play == nil {
		return gorm.ErrRecordNotFound	
	}
	return mysql.NewPlayDao().UpdatePlay(int(play_id), play_name, play_description, play_start, play_end, play_price)
}
