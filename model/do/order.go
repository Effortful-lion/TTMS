package do

import "time"

type Order struct {
	OrderID   int       `gorm:"primaryKey;autoIncrement"` // 订单ID，主键，自增
	OrderTime time.Time `gorm:"type:datetime"`            // 订单时间（数据库datetime类型）

	// 依赖
	UserInfo   UserInfo
	UserInfoID int64 `gorm:"foreignKey:UserInfoID;references:UserInfoID;index;not null"` // 用户ID，外键
}