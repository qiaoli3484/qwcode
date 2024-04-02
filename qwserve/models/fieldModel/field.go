package fieldModel

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"qwserve/models"
	"time"
)

type Field struct {
	ID        uint   `json:"id" gorm:"primarykey"`
	Alias     string `json:"alias" gorm:"type:varchar(10);"`    //别称
	Name      string `json:"name" gorm:"type:varchar(50);"`     //字段名
	DataType  string `json:"dataType" gorm:"type:varchar(20);"` //字段类型
	SecLevel  int    `json:"secLevel" gorm:"type:tinyint;"`     //保密级别
	CbVisible bool   `json:"cbVisible" gorm:"type:tinyint;"`    //是否隐藏 编辑隐藏
	Cbsearch  bool   `json:"cbsearch" gorm:"type:tinyint;"`     //是搜索
	Explain   string `json:"explain" gorm:"type:varchar(255);"` //提示
	Comment   string `json:"comment" gorm:"type:varchar(255);"` //备注

	StrLen     int    `json:"strLen" gorm:"type:int;"`              //文本长度
	DefaultVal string `json:"defaultVal" gorm:"type:varchar(255);"` //默认值
	Required   bool   `json:"required" gorm:"type:tinyint;"`        //必填字段
	EnDistinct bool   `json:"enDistinct" gorm:"type:tinyint;"`      //不允许重复

	EnlinkTo bool                              `json:"enlinkTo" gorm:"tinyint;"`
	LinkTo   `json:"linkTo" gorm:"type:TEXT;"` //链接到

	EnAutoComplete bool `json:"enAutoComplete" gorm:"tinyint;"` //输入提示
	AutoComplete   `json:"autoComplete" gorm:"type:TEXT;"`

	EnFormula   bool   `json:"enFormula" gorm:"tinyint;"`   //使用公式
	Formula     string `json:"formula" gorm:"type:text;"`   //放公式
	FormulaMode int    `json:"formulaMode" gorm:"tinyint;"` //1写2读

	Validate   string `json:"validate" gorm:"type:varchar(255);"`   //验证
	ValExplain string `json:"valexplain" gorm:"type:varchar(255);"` //验证提示

	Style `json:"style" gorm:"type:TEXT;"` //链接到

	CreatedAt time.Time
	UpdatedAt time.Time
}

type Style struct {
	Width string `json:"width"` //宽度
	Fixed string `json:"fixed"` //固定
	Align string `json:"align"` //对齐方式
	Hide  bool   `json:"hide"`  //隐藏
	Tonj  string `json:"tonj"`  //统计

	EnDisplay bool             `json:"enDisplay"` //编辑隐藏
	Filters   []*models.Filter `json:"filters"`   //类型
}

//链接
type LinkTo struct {
	AcTemplet int64            `json:"acTemplet"` //类型
	Rb        string           `json:"rb"`
	Target    string           `json:"Target"`   //类型
	Enfilter  bool             `json:"enfilter"` //类型
	Filter    string           `json:"filter"`   //类型
	Filters   []*models.Filter `json:"filters"`  //类型
	Host      string           `json:"host"`     //类型
	Alias     string           `json:"alias"`    //类型
}

type AutoComplete struct {
	AcTemplet int64                  `json:"acTemplet"` //类型
	Filter    string                 `json:"filter"`    //类型
	Enfilter  bool                   `json:"enfilter"`  //类型
	Filters   []*models.Filter       `json:"filters"`   //类型
	Result    map[string]interface{} `json:"result"`    //类型
}

func (j *Style) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}
	result := Style{}
	err := json.Unmarshal(bytes, &result)
	*j = Style(result)
	return err
}

// 实现 driver.Valuer 接口，Value 返回 json value
func (j Style) Value() (driver.Value, error) {
	return json.Marshal(j)
}

func (j *LinkTo) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}
	result := LinkTo{}
	err := json.Unmarshal(bytes, &result)
	*j = LinkTo(result)
	return err
}

// 实现 driver.Valuer 接口，Value 返回 json value
func (j LinkTo) Value() (driver.Value, error) {
	return json.Marshal(j)
}

func (j *AutoComplete) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}
	result := AutoComplete{}
	err := json.Unmarshal(bytes, &result)
	*j = AutoComplete(result)
	return err
}

// 实现 driver.Valuer 接口，Value 返回 json value
func (j AutoComplete) Value() (driver.Value, error) {
	return json.Marshal(j)
}
