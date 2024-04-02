package fieldModel

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

// 图片
type Photo struct {
	Field
	Extend *Pextend `json:"extend" gorm:"type:TEXT;"`
}
type Pextend struct {
	//待添加 上传数量 限制大小 限制类型
	Limit  int     `json:"limit"`  //允许上传文件的最大数量
	Size   float32 `json:"size"`   //上传尺寸
	Accept string  `json:"accept"` //接受上传的文件类型
}

func (j Pextend) Value() (driver.Value, error) {
	return json.Marshal(j)
}

func (j *Pextend) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}
	result := Pextend{}
	err := json.Unmarshal(bytes, &result)
	*j = Pextend(result)
	return err
}
