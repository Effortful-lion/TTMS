package do



type Resource struct {
	ResourceID   int64    		`gorm:"column:resource_id;type:int;primaryKey;autoIncrement"` // 资源ID，主键，自增
	ResourceName string   		`gorm:"column:resource_name;type:varchar(100);unique"`     // 资源名称，非空且唯一
	ResourceURL  string   		`gorm:"column:resource_url;type:varchar(100);unique"`     // 资源URL，非空且唯一
}

// TableName 返回表名
func (Resource) TableName() string {
	return "resource"
}