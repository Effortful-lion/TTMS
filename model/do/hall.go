package do


//----------------------------演出厅---------------------------
type Hall struct {
	HallID    int64  `gorm:"column:hall_id;type:int;primaryKey;autoIncrement"`   // 演出厅ID，主键，自增
	HallName  string `gorm:"column:hall_name;type:varchar(100);not null;unique"` // 演出厅名称（唯一约束）
	HallRow   int    `gorm:"column:hall_row;type:int;not null"`                  // 行数
	HallCol   int    `gorm:"column:hall_col;type:int;not null"`                  // 列数
	HallTotal int    `gorm:"column:hall_total;type:int;not null"`                // 座位总数（应等于 HallRow*HallCol）
}

// TableName 定义表名
func (Hall) TableName() string {
	return "hall" // 对应数据库中的表名
}