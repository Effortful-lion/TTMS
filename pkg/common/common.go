package common

import (
	"TTMS/model/dto"
	"fmt"
	"math"
	"strconv"
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