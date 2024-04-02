package fieldModel

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

// 变量
type Variable struct {
	Field
	Extend *Vextend `json:"extend" gorm:"type:TEXT;"`
}

// 变量
type Vextend struct {
	Typ    string `json:"typ"`    //变量类型
	Format string `json:"format"` //时间格式
	Unit   string `json:"unit"`   //机构

	EnUpdate bool `json:"enUpdate"` //编辑时更新变量值
}

func (j Vextend) Value() (driver.Value, error) {
	return json.Marshal(j)
}

func (j *Vextend) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}
	result := Vextend{}
	err := json.Unmarshal(bytes, &result)
	*j = Vextend(result)
	return err
}
