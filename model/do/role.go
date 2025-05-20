package do

type Role struct {
	RoleID   int64    		`gorm:"column:role_id;type:int;primaryKey;autoIncrement"` // 角色ID，主键，自增
	RoleName string   		`gorm:"column:role_name;type:varchar(100);not null;unique"`     // 角色名称，非空且唯一
}

// TableName 返回表名
func (Role) TableName() string {
	return "role"
}