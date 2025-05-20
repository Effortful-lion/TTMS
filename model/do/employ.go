package do


type Employ struct {
	EmployID   int64    		`gorm:"column:employ_id;type:int;primaryKey;autoIncrement"` // 员工ID，主键，自增
	EmployName string   		`gorm:"column:employ_name;type:varchar(100);not null;unique"`     // 员工姓名，非空且唯一
	EmployPassword string 	`gorm:"column:employ_password;type:varchar(100);not null"` // 员工密码，非空
}

// TableName 返回表名
func (Employ) TableName() string {
	return "employ"
}