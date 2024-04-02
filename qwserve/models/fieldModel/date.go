package fieldModel

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

// 日期
type Tdate struct {
	Field
	Extend *Dextend `json:"extend" gorm:"type:TEXT;"`
}

type Dextend struct {
	Format string `json:"format"` //日期格式
}

func (j Dextend) Value() (driver.Value, error) {
	return json.Marshal(j)
}

func (j *Dextend) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}
	result := Dextend{}
	err := json.Unmarshal(bytes, &result)
	*j = Dextend(result)
	return err
}
