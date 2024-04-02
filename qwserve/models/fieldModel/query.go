package fieldModel

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"qwserve/models"
)

// 查询字段
type Query struct {
	Field
	Extend *Qextend `json:"extend" gorm:"type:TEXT;"`
}

type Qextend struct {
	AcTemplet int64                  `json:"acTemplet"` //类型
	Enfilter  bool                   `json:"enfilter"`  //类型
	Filter    string                 `json:"filter"`    //类型
	Filters   []*models.Filter       `json:"filters"`   //类型
	Ra        string                 `json:"ra"`        //1 取值 2 合计 3计数 count
	Result    map[string]interface{} `json:"result"`    //类型
}

func (j Qextend) Value() (driver.Value, error) {
	return json.Marshal(j)
}

func (j *Qextend) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}
	result := Qextend{}
	err := json.Unmarshal(bytes, &result)
	*j = Qextend(result)
	return err
}
