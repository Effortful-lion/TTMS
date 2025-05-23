package mysql

import (
	"TTMS/model/do"
	"errors"
)

type PlayDao struct {
}

func NewPlayDao() *PlayDao {
	return &PlayDao{}
}

func (ud *PlayDao) SelectNamesByIDs(play_ids []int64) ([]string, error) {
	names := make([]string, len(play_ids))
	for i := range play_ids {
		play_id := play_ids[i]
		// 根据id获得name
		name, err := ud.SelectNameByID(play_id)
		if err != nil{return []string{}, err}
		names[i] = name
	}
	return names, nil
}

func (ud *PlayDao) SelectNameByID(play_id int64) (string, error) {
	var play do.Play
	err := DB.First(&play, play_id).Error
	if err != nil {
		if err.Error() == "record not found" {
			return "", errors.New("SelectNameByID record not found")	
		}
		return "", err
	}
	return play.PlayName, nil
}

func (ud *PlayDao) InsertPlay(play_name, play_description string, play_duration int) error {
	play := &do.Play{
		PlayName:        play_name,
		PlayDescription: play_description,
		PlayDuration:    play_duration,
	}
	return DB.Create(play).Error
}

func (ud *PlayDao)DeletePlay(play_id int64) error {
	var play *do.Play
	err := DB.Where("play_id =?", play_id).Delete(&play).Error
	if err!= nil {
		return err
	}
	// 剧目删除，对应的plan也删除
	err = NewPlanDao().DeletePlanByPlayID(play_id)
	if err!= nil {
		return err
	}
	return nil
}

func (*PlayDao)UpdatePlay(play_id int,play_name, play_description string, play_duration int) error {
	play := &do.Play{
		PlayID: int64(play_id),
		PlayName: play_name,
		PlayDescription: play_description,
		PlayDuration: play_duration,
	}
	return DB.Save(play).Error
}

func (*PlayDao) SelectPlayByID(play_id int64) (*do.Play, error) {
	var play do.Play
	err := DB.Where("play_id = ?", play_id).First(&play).Error	
	if err!= nil {
		if err.Error() == "record not found" {
			return nil, errors.New("SelectPlayByID record not found")
		}
		return nil, err	
	}
	return &play, nil
}

func (*PlayDao)SelectAllPlay() ([]*do.Play, error) {
	var plays []do.Play
	err := DB.Find(&plays).Error
    if err != nil {
        return nil, err
    }
    // 将非指针切片转换为指针切片
    playPtrs := make([]*do.Play, len(plays))
    for i := range plays {
        playPtrs[i] = &plays[i]
    }
    return playPtrs, nil
}

