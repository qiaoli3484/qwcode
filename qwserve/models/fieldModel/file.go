package fieldModel

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

// 文件
type File struct {
	Field
	Extend *Fextend `json:"extend" gorm:"type:TEXT;"`
}

type Fextend struct {
	Accept string `json:"accept"` //限制类型  doc,docx,jpg,jpeg,png,xlsx,pdf

	Limit int     `json:"limit"` //上传数量
	Size  float32 `json:"size"`  //限制大小
}

func (j Fextend) Value() (driver.Value, error) {
	return json.Marshal(j)
}

func (j *Fextend) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}
	result := Fextend{}
	err := json.Unmarshal(bytes, &result)
	*j = Fextend(result)
	return err
}
