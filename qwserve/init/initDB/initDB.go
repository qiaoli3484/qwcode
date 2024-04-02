package initDB

import (
	"qwserve/models"
	"qwserve/models/fieldModel"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitMysql() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&loc=Local&parseTime=True" //&parseTime=True
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB = d
	sqlDB, err := d.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(500)
	sqlDB.SetConnMaxLifetime(time.Hour)

	// 迁移 schema
	d.AutoMigrate(&models.WebComponents{})
	// 根据 User 的字段创建 `deleted_users` 表
	d.Table("fields").AutoMigrate(&fieldModel.Ttext{})
}
