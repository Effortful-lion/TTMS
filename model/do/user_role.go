package do


type UserRole struct {
	UserRoleID   int64    		`gorm:"column:user_role_id;type:int;primaryKey;autoIncrement"` // 用户角色ID，主键，自增
	EmployID     int64			`gorm:"column:employ_id;type:int;not null"`     // 员工ID，非空
	RoleID       int64			`gorm:"column:role_id;type:int;not null"`     // 角色ID，非空
}

// TableName 返回表名
func (UserRole) TableName() string {
	return "user_role"
}