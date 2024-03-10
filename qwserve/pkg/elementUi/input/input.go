package input

import (
	_ "embed"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"qwserve/pkg/elementUi/base"
	"qwserve/pkg/elementUi/util"
)

//go:embed input.vue
var Inputjs string // 当前目录，解析为string类型

type ElInput struct {
	base.Base
	size         string `att:"size" default:"small"`
	typ          string `att:"type" default:"text"`        //text，textarea
	maxlength    int    `att:"maxlength" default:"150"`    //原生属性，最大输入长度
	minlength    int    `att:"minlength" default:"0"`      //原生属性，最小输入长度
	placeholder  string `att:"placeholder" default:"true"` //输入框占位文本
	clearable    bool   `att:"clearable" default:"true"`
	showpassword bool   `att:"show-password" default:"false"`
	readonly     bool   `att:"readonly" default:"false"`
}

type Option interface {
	Configure(*ElInput) error
}
type OptionFn func(*ElInput) error

func (o OptionFn) Configure(rl *ElInput) error {
	return o(rl)
}

func New(options ...Option) *ElInput {
	rl := &ElInput{}
	for _, opt := range options {
		opt.Configure(rl)
	}
	return rl
}

func WithType(typ string) Option {
	return OptionFn(func(rl *ElInput) error {
		rl.typ = typ
		return nil
	})
}

func WithValue(typ string) Option {
	return OptionFn(func(rl *ElInput) error {
		rl.typ = typ
		return nil
	})
}

func (el *ElInput) View(alias string) string {
	if alias != "" {
		alias = alias + "."
	}
	valueof := reflect.Indirect(reflect.ValueOf(el))
	typeof := valueof.Type()
	num := typeof.NumField()
	arr := make([]string, 0)

	for i := 0; i < num; i++ {
		sf := typeof.Field(i)
		sv := valueof.Field(i)
		key := sf.Tag.Get("att")
		def := sf.Tag.Get("default")
		if key == "" {
			continue
		}

		if sf.Type.Kind() == reflect.String {
			if sv.String() == "" {
				arr = append(arr, key+"='"+def+"'")
				continue
			}
			arr = append(arr, key+"='"+sv.String()+"'")

		} else if sf.Type.Kind() == reflect.Int {
			if sv.Int() == 0 {
				arr = append(arr, key+":='"+def+"'")
				continue
			}
			arr = append(arr, key+":='"+strconv.Itoa(int(sv.Int()))+"'")
		} else if sf.Type.Kind() == reflect.Bool {
			bb, _ := strconv.ParseBool(def)

			if sv.Bool() || bb {
				arr = append(arr, key)
				continue
			}
			arr = append(arr, key+":='false'")
		}

	}
	return "<el-input v-model='" + alias + el.Getkey() + "' " + strings.Join(arr, " ") + "></el-input>"
}

func (el *ElInput) Data() string     { return el.Getkey() + ":\"" + "" + "\"" }
func (el *ElInput) Method() string   { return "" }
func (el *ElInput) Init() string     { return "" }
func (el *ElInput) Computed() string { return "" }
func (el *ElInput) Component() map[string]string {
	return map[string]string{"inputjs": Inputjs}
}

func (el *ElInput) GetRule() string {
	s1 := `%s: [
	{ required: true, message: '请输入%s', trigger: 'blur' }
  ],`
	return fmt.Sprintf(s1, el.Getkey(), el.GetName())
}

type ElNumber struct {
	base.Base

	size        string `att:"size" default:"small"`
	precision   int    `att:"precision" default:"0"`
	controls    bool   `att:"controls" default:"false"`
	readonly    bool   `att:"readonly" default:"false"`
	placeholder string `att:"placeholder" default:"输入"` //输入框占位文本
	clearable   bool   `att:"clearable" default:"true"`
	minwidth    string `style:"min-width" default:"60px!important"`
	width       string `style:"width" default:"100%!important"`
}

func NewNumber() *ElNumber {
	rl := &ElNumber{}
	/*for _, opt := range options {
		opt.Configure(rl)
	}*/
	return rl
}

func (el *ElNumber) View(alias string) string {
	if alias != "" {
		alias = alias + "."
	}

	return "<el-input-number v-model=\"" + alias + el.Getkey() + "\" " + util.AttrRef(el) + "/>"
}
func (el *ElNumber) SetName(val string) *ElNumber {
	el.Base.SetName(val)
	return el
}

func (el *ElNumber) Setkey(val string) *ElNumber {
	el.Base.Setkey(val)
	return el
}

func (el *ElNumber) Data() string     { return el.Getkey() + ":" + "0" + "" }
func (el *ElNumber) Method() string   { return "" }
func (el *ElNumber) Init() string     { return "" }
func (el *ElNumber) Computed() string { return "" }
func (el *ElNumber) Component() map[string]string {
	return nil
}

func (el *ElNumber) GetRule() string {
	s1 := `%s: [
	{ required: true, message: '请输入%s', trigger: ['blur', 'change'] }
  ],`
	return fmt.Sprintf(s1, el.Getkey(), el.GetName())
}
