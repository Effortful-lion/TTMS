package do


//----------------------------演出厅---------------------------
type Hall struct {
	HallID   int64  `gorm:"column:hall_id;type:int;primaryKey;autoIncrement"`  // 演出厅ID，主键，自增
	HallName string `gorm:"column:hall_name;type:varchar(100);not null"`        // 演出厅名称，唯一
	HallRow  int    `gorm:"column:hall_row;type:int;not null"`                  // 行数
	HallCol  int    `gorm:"column:hall_col;type:int;not null"`                  // 列数
	HallTotal int    `gorm:"column:hall_total;type:int;not null"`               // 座位总数

	// 依赖 演出计划 一对多： 一个演出计划 -- 多个演出厅
}