package fieldModel

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

// Text: "文本"
type Ttext struct {
	Field
	Extend *Textend `json:"extend" gorm:"type:TEXT;"`
}

type Textend struct {
	EnEdit bool `json:"enEdit"` //使用编辑器
}

func (j Textend) Value() (driver.Value, error) {
	return json.Marshal(j)
}

func (j *Textend) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}
	result := Textend{}
	err := json.Unmarshal(bytes, &result)
	*j = Textend(result)
	return err
}
