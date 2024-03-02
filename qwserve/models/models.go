package models

import "gorm.io/gorm"

//网页组件
type WebComponents struct {
	gorm.Model
	Name  string `gorm:"type:varchar(65);default:'';comment:名称;"`
	Code  string `gorm:"type:varchar(65);default:'';comment:代码;"`
	State uint8  `gorm:"type:tinyint;default:0;comment:状态;"`
	Price uint
}

// TableName 会将 User 的表名重写为 `profiles`
/*func (Product) TableName() string {
	return "WebComponents"
}*/
