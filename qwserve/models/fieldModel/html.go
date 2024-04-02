package fieldModel

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

// 文件
type Html struct {
	Field
	Extend *Hextend `json:"extend" gorm:"type:TEXT;"`
}

type Hextend struct {
	View   string `json:"view"`
	Script string `json:"script"`
}

func (j Hextend) Value() (driver.Value, error) {
	return json.Marshal(j)
}

func (j *Hextend) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}
	result := Hextend{}
	err := json.Unmarshal(bytes, &result)
	*j = Hextend(result)
	return err
}
