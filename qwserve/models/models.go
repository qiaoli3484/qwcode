package models

import (
	"time"

	"gorm.io/gorm"
)

//网页组件
type WebComponents struct {
	ID        uint   `json:"id" gorm:"primarykey"`
	Name      string `json:"name" gorm:"type:varchar(65);default:'';comment:名称;"`
	Code      string `json:"code" gorm:"type:TEXT;comment:代码;"`
	State     uint8  `json:"state" gorm:"type:tinyint;default:0;comment:状态;"`
	Comment   string `json:"comment" gorm:"type:varchar(255);comment:代码;"`
	Price     uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// TableName 会将 User 的表名重写为 `profiles`
/*func (Product) TableName() string {
	return "WebComponents"
}*/
