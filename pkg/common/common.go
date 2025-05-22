package common

import (
	"strconv"
	"time"

)

func ParseStringTime(timeStr string) (time.Time) {
	if timeStr == "" {
		return time.Now()	
	}
	parsedTime, err := time.ParseInLocation("2006-01-02 15:04:05", timeStr, time.Local)
	if err != nil {
		return time.Now()
	}
	return parsedTime
}

// 将时间转换为字符串
func ParseTimeToString(t time.Time) string {
    return t.Format("2006-01-02 15:04:05")
}

func ParseStringTimeToTimeStamp(timeStr string) (int64) {
	if timeStr == "" {
		return time.Now().Unix()
	}
	parsedTime, err := time.ParseInLocation("2006-01-02 15:04:05", timeStr, time.Local)
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