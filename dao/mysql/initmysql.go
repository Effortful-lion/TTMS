package mysql

// 这里是关于mysql的读取配置，初始化的方法等

import (
	"TTMS/config"
	"TTMS/model/do"
	"fmt"
	"log"

	"gorm.io/driver/mysql" // gorm内置的mysql驱动
	"gorm.io/gorm"
)

var DB *gorm.DB // 这里是一个全局变量，用来存储mysql的连接实例

func InitMysql() (err error) {
	cfg := config.Conf.MysqlConfig
	// 连接数据库: 数据源名称
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Dbname)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: nil,
	})
	if err != nil {
		return err
	}

	// 连接成功后，设置其他配置：
	sqlDB, err := DB.DB()
	if err != nil {
		return
	}
	// 设置最大连接数
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
	// 设置最大空闲连接数
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
	log.Println("mysql连接成功！")
	MigrateTables()
	// 初始化角色信息
	InitRole()
	SyncData()
	return
}

func InitRole() {
	// 初始化角色信息
	roles := []do.Role{
		{RoleName: "admin"},
		{RoleName: "staff"},
		{RoleName: "manager"},
		{RoleName: "ticketor"},
		{RoleName: "finance"},
		{RoleName: "account"},
		{RoleName: "customer"},
	}

	for _, role := range roles {
		var count int64
		// 检查角色是否已存在
		DB.Model(&do.Role{}).Where("role_name = ?", role.RoleName).Count(&count)
		if count == 0 {
			// 角色不存在，插入数据
			if err := DB.Create(&role).Error; err != nil {
				log.Fatalf("Failed to insert role %s: %v", role.RoleName, err)
			}
		}
	}
}

// 自动迁移表结构
func MigrateTables() error {
	return DB.AutoMigrate(
		// 在这里添加需要迁移的表结构
		do.Customer{},
		do.Employ{},
		do.Resource{},
		do.Role{},
		do.RoleResource{},
		do.UserRole{},
		do.Play{},
		do.Hall{},
		do.Plan{},
		do.Seat{},
		do.Ticket{},
		do.Sale{},
	)	
}