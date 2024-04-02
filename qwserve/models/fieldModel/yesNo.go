package fieldModel

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

// 是否
type YesNo struct {
	Field

	Extend *Yextend `json:"extend" gorm:"type:TEXT;"`
}

type Yextend struct {
}

func (j Yextend) Value() (driver.Value, error) {
	return json.Marshal(j)
}

func (j *Yextend) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}
	result := Yextend{}
	err := json.Unmarshal(bytes, &result)
	*j = Yextend(result)
	return err
}
