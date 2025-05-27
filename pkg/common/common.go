package common

import (
	"TTMS/model/dto"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
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

// ParseStringTime 解析各种常见时间格式，优先支持带时区的格式
func ParseStringTime(timeStr string) (time.Time) {
	if timeStr == "" {
		return time.Now()
	}

	layouts := []string{
		time.RFC3339,                   // "2025-05-27T14:30:00+08:00"
		"2006-01-02T15:04:05Z07:00",   // 同 RFC3339
		"2006-01-02 15:04:05 -0700",   // 带时区偏移
		"2006-01-02 15:04:05 MST",     // 带时区缩写
		"2006-01-02 15:04:05",         // 仅日期时间
		"2006/01/02 15:04:05",         // 斜杠分隔
		"2006-01-02",                  // 仅日期
	}

	var parsedTime time.Time
	var err error

	for _, layout := range layouts {
		parsedTime, err = time.Parse(layout, timeStr)
		if err == nil {
			return parsedTime
		}
	}

	// 如果都失败，尝试拆分解析（兼容老逻辑）
	parts := strings.Split(timeStr, " ")
	if len(parts) >= 2 {
		dateTimeStr := parts[0] + " " + parts[1]
		parsedTime, err = time.Parse("2006-01-02 15:04:05", dateTimeStr)
		if err == nil {
			return parsedTime
		}
	}

	return time.Time{}
}

// 将时间转换为字符串
func ParseTimeToString(t time.Time) string {
    return t.Format("2006-01-02 15:04:05")
}

// ParseStringTimeToTimeStamp 将字符串时间转为 Unix 时间戳（秒）
func ParseStringTimeToTimeStamp(timeStr string) (int64) {
	t := ParseStringTime(timeStr)
	return t.Unix()
}

func ParseStringToInt64(str string) (int64, error) {
	if str == "" {
		return 0, nil
	}
	str_to_id, err := strconv.Atoi(str)
	return int64(str_to_id), err
}

// 结构体排名
func SortStructByField(data []dto.TicketCountResp, field string) {
	// 根据 field 字段对 data 进行排序
	for i := 0; i < len(data)-1; i++ {
		for j := i + 1; j < len(data); j++ {
			if data[i].TotalMoney < data[j].TotalMoney {
				data[i], data[j] = data[j], data[i]
			}
		}
	}
}

// 放入context中
func SetContext(c *gin.Context, key string, value interface{}) {
	c.Set(key, value)
}

// 从context中取出
func GetContext(c *gin.Context, key string) (interface{}, bool) {
	value, exists := c.Get(key)
	return value, exists	
}

// 比较时间前后// 比较时间前后
// 前面的早于后面的返回 true
func TimeAfter(start, end string) bool {
    return ParseStringTimeToTimeStamp(start) < ParseStringTimeToTimeStamp(end)
}