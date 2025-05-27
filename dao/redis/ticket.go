package redis

import (
	"context"
	"fmt"
	"time"
)

const TicketExpiredTime	= 5 * time.Minute

func InsertTicketOther(key string,customerID, planID, seatID int64, expire time.Duration) (err error){
	// 构建一个 hash 存储
	// key 为订单号，其他为订单信息
	values := map[string]interface{}{
		"customerID": customerID,
		"planID": planID,
		"seatID": seatID,
	}
	
	// 使用 HMset 设置哈希表字段值
	if err := Rdb.HMSet(context.Background(), key, values).Err(); err != nil {
		return err
	}

	ctx := context.Background()
	// 使用Expire设置过期时间
    if err := Rdb.Expire(ctx, key, TicketExpiredTime).Err(); err != nil {
        return err
    }

	return
}

func GetTicketOther(key string) (map[string]string, error) {
	return Rdb.HGetAll(context.Background(), key).Result()
}

func InsertTicket(customerID, planID, seatID int64, customerName string, price float64, expire time.Time, playID int64, authID int8) error {
    key := fmt.Sprintf("ticket:%d_%d_%d", customerID, planID, seatID)
    
    values := map[string]interface{}{
        "customerID":   fmt.Sprintf("%d", customerID),
        "planID":       fmt.Sprintf("%d", planID),
        "seatID":       fmt.Sprintf("%d", seatID),
        "customerName": customerName,
        "price":        fmt.Sprintf("%f", price),
        "expire":       expire.String(), // 把Duration转换为字符串
        "playID":       fmt.Sprintf("%d", playID),
        "authID":       fmt.Sprintf("%d", authID),
    }

    // 使用HMSet设置哈希表的字段值
    ctx := context.Background()
    cmd := Rdb.HMSet(ctx, key, values)
    if cmd.Err() != nil {
        return cmd.Err()
    }

    // 使用Expire设置过期时间
    if err := Rdb.Expire(ctx, key, TicketExpiredTime).Err(); err != nil {
        return err
    }

    return nil
}
func GetTicket(customerID, planID, seatID int64) (map[string]string, error) {
	key := fmt.Sprintf("ticket:%d_%d_%d", customerID, planID, seatID)
	fmt.Println(key)
	return Rdb.HGetAll(context.Background(), key).Result()
}
