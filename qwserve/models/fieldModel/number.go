package fieldModel

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

// 数字
type Number struct {
	Field
	Extend *Nextend `json:"extend" gorm:"type:TEXT;"`
}

type Nextend struct {
	Currency    string `json:"Currency"`    //货币符号
	Unit        string `json:"Unit"`        //单位
	BitsMode    string `json:"bitsMode"`    //位数
	Bits        int    `json:"bits"`        //位数
	Special     bool   `json:"Special"`     //显示
	DdlSpecial  string `json:"ddlSpecial"`  //方式
	SciNotation bool   `json:"SciNotation"` //科学
	Custom      bool   `json:"Custom"`
	CbCustom    string `json:"cbCustom"` //自定义
}

func (j Nextend) Value() (driver.Value, error) {
	return json.Marshal(j)
}

func (j *Nextend) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}
	result := Nextend{}
	err := json.Unmarshal(bytes, &result)
	*j = Nextend(result)
	return err
}
