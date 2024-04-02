package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
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
	Param     CompParam `json:"param" gorm:"type:TEXT;comment:代码;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// TableName 会将 User 的表名重写为 `profiles`
/*func (Product) TableName() string {
	return "WebComponents"
}*/

//组件参数
type CompParam struct {
	Binds  []interface{} `json:"binds"`
	Styles []interface{} `json:"styles"`
}

//{title:'',param:{default:'',params:[]}
// 实现 sql.Scanner 接口，Scan 将 value 扫描至 Jsonb
func (j *CompParam) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}
	result := CompParam{}
	err := json.Unmarshal(bytes, &result)
	*j = CompParam(result)
	return err
}

// 实现 driver.Valuer 接口，Value 返回 json value
func (j CompParam) Value() (driver.Value, error) {
	fmt.Println(j, "Value")
	return json.Marshal(j)
}
