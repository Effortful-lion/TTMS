package do

import "time"

// 售票信息
type Sale struct {
	SaleID int64 `gorm:"column:sale_id;type:int;primaryKey;autoIncrement"`		// 售票ID
	EmployID int64 `gorm:"column:employ_id;type:int;not null"`					// 售出人ID
	CustomerID int64 `gorm:"column:customer_id;type:int;not null"`				// 买票人ID
	SaleTime time.Time `gorm:"column:sale_time;type:datetime;not null"`			// 售票时间
	SalePrice float64 `gorm:"column:sale_price;type:float(10,2);not null"`		// 售票价格
	TicketID int64 `gorm:"column:ticket_id;type:int;not null"`					// 购买了哪一张票
	SaleStatus int8 `gorm:"column:sale_status;type:tinyint;not null"`			// 售票状态
}

const (
	// 售卖状态：结合 支付状态和票的状态
	// 未支付、已支付、已退票、已取消 （已取消：已支付+已核销、已退票、未支付+订单过期、已支付+未核销+票过期）
	SaleStatusUnPay = iota			// 未支付（和票无关）
	SaleStatusPayed					// 已支付（有票：未核销、已核销、已取消）
	SaleStatusRefund				// 已退票（和票无关）
	SaleStatusCancel				// 已取消（和票无关）
)

func (Sale) TableName() string {
	return "sale"
}