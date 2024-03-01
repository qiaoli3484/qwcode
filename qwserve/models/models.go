package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Code  string `gorm:"type:varchar(65);default:'';comment:测试;"`
	Price uint
}
