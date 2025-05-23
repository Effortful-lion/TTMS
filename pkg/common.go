package pkg

import (
	"strconv"
	"time"
)

func ParseStringTime(timeStr string) (time.Time) {
	if timeStr == "" {
		return time.Now()	
	}
	parsedTime, err := time.Parse("2006-01-02 15:04:05", timeStr)
	if err != nil {
		return time.Now()
	}
	return parsedTime
}

func ParseStringTimeToTimeStamp(timeStr string) (int64) {
	if timeStr == "" {
		return time.Now().Unix()
	}
	parsedTime, err := time.Parse("2006-01-02 15:04:05", timeStr)
	if err!= nil {
		return time.Now().Unix()
	}
	return parsedTime.Unix()
}

func ParseStringToInt64(str string) (int64, error) {
	if str == "" {
		return 0, nil
	}
	str_to_id, err := strconv.Atoi(str)
	return int64(str_to_id), err
}