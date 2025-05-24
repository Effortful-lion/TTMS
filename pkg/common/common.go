package common

import (
	"fmt"
	"math"
	"strconv"
	"time"
)

func ToTwoDecimal(num float64) float64 {
    return math.Round(num*100) / 100
}

// 计算百分比 string
// 计算百分比并格式化为字符串（如 "25.00%"）
func CalculatePercentageString(part, total float64) string {
    if total == 0 {
        return "0.00%"
    }
    return fmt.Sprintf("%.2f%%", (part/total)*100)
}

// 计算百分比 float64
func CalculatePercentageFloat(part, total float64) float64 {
	// 如果分母为0，返回0.00
	if total == 0 {
		return 0.00
	}else{
		// 返回两位小数的百分比
		return ToTwoDecimal(part*100/total)
	}
}

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