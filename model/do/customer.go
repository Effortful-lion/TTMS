package do



type Customer struct {
	CustomerID   int64    		`gorm:"column:customer_id;type:int;primaryKey;autoIncrement"` // 客户ID，主键，自增
	CustomerName string   		`gorm:"column:customer_name;type:varchar(100);not null;unique"`     // 客户姓名，非空且唯一
	CustomerPassword string 	`gorm:"column:customer_password;type:varchar(100);not null"` // 客户密码，非空
}

// TableName 返回表名
func (Customer) TableName() string {
	return "customer"
}

