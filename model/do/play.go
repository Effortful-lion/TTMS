package do



type Play struct {
	PlayID          int64     `gorm:"primaryKey;autoIncrement"`
	PlayName        string    `gorm:"type:varchar(255);not null;unique"` // 剧目名称（唯一约束）
	PlayDescription string    `gorm:"type:varchar(255);not null"`        // 剧目描述
	PlayDuration    int 	  `gorm:"type:int;not null"`		 		 // 剧目时长（单位：分钟），使用 int 类型
}

// TableName 自定义表名
func (Play) TableName() string {
	return "play" // 指定表名为 "play"
}