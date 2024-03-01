package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name  string `gorm:"type:varchar(65);default:'';comment:名称;"`
	Code  string `gorm:"type:varchar(65);default:'';comment:代码;"`
	State uint8  `gorm:"type:tinyint;default:0;comment:状态;"`
	Price uint
}
