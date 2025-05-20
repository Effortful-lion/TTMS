package do



type RoleResource struct {
	RoleResourceID   int64    		`gorm:"column:role_resource_id;type:int;primaryKey;autoIncrement"` // 角色资源ID，主键，自增
	RoleID   int64    		`gorm:"column:role_id;type:int;not null"` // 角色ID，非空
	ResourceID   int64    		`gorm:"column:resource_id;type:int;not null"` // 资源ID，非空
}

// TableName 返回表名
func (RoleResource) TableName() string {
	return "role_resource"
}