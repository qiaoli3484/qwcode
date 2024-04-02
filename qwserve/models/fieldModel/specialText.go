package fieldModel

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

// 特殊文本
type SpecialText struct {
	Field
	Extend *Spextend `json:"extend" gorm:"type:TEXT;"`
}

// 特殊文本
type Spextend struct {
	Typ string `json:"typ"`

	Users    []string `json:"users"`    //用户
	Roles    []int64  `json:"roles"`    //角色
	Companys []int64  `json:"companys"` //机构
	Exclude  []string `json:"exclude"`  //排除
	Rector   bool     `json:"rector"`   //负责人
	Directly bool     `json:"directly"` //直属人员
	Indirect bool     `json:"indirect"` //非直属人员
}

func (j Spextend) Value() (driver.Value, error) {
	return json.Marshal(j)
}

func (j *Spextend) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}
	result := Spextend{}
	err := json.Unmarshal(bytes, &result)
	*j = Spextend(result)
	return err
}
