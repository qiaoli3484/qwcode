package fieldModel

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

// 自动编号
type AutoNumber struct {
	Field
	Extend *Aextend `json:"extend"  gorm:"type:TEXT;"`
}

type Aextend struct {
	EnPrefix bool   `json:"enPrefix"` //前缀
	Prefix   string `json:"Prefix"`
	EnRandom bool   `json:"enRandom"` //随机码
	EnTime   bool   `json:"enTime"`   //日期
	Time     string `json:"Time"`     //日期格式
	EnSerial bool   `json:"enSerial"` //流水号
	Bits     int    `json:"bits"`     //位数
	Val      int    `json:"val"`
	Add      int    `json:"add"`
	Cycle    string `json:"cycle"`
	EnSuffix bool   `json:"enSuffix"` //后缀
	Suffix   string `json:"Suffix"`   //后缀编号
}

func (j Aextend) Value() (driver.Value, error) {
	return json.Marshal(j)
}

func (j *Aextend) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}
	result := Aextend{}
	err := json.Unmarshal(bytes, &result)
	*j = Aextend(result)
	return err
}
