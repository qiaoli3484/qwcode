package fieldModel

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

// 单选 或 多选
type SingleList struct {
	Field
	Extend *Sextend `json:"extend" gorm:"type:TEXT;"`
}

type Sextend struct {
	Multiple bool `json:"multiple"`

	List    string `json:"list"`
	EnColor bool   `json:"enColor"`
}

func (j Sextend) Value() (driver.Value, error) {
	return json.Marshal(j)
}

func (j *Sextend) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}
	result := Sextend{}
	err := json.Unmarshal(bytes, &result)
	*j = Sextend(result)
	return err
}
