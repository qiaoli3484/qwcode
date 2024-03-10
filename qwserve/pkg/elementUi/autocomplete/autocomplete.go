package autocomplete

import (
	_ "embed"
	"reflect"
	"simpleTool/pkg/vue/elementUi/base"
	"strconv"
	"strings"
)

//go:embed autocomplete.js
var autocompletejs string // 当前目录，解析为string类型

type Autocomplete struct {
	base.Base
	size         string `att:"size" default:"mini"`
	typ          string `att:"type" default:"text"`        //text，textarea
	maxlength    int    `att:"maxlength" default:"150"`    //原生属性，最大输入长度
	minlength    int    `att:"minlength" default:"0"`      //原生属性，最小输入长度
	placeholder  string `att:"placeholder" default:"true"` //输入框占位文本
	clearable    bool   `att:"clearable" default:"true"`
	showpassword bool   `att:"show-password" default:"true"`
	disabled     bool   `att:"disabled"`
}

type Option interface {
	Configure(*Autocomplete) error
}
type OptionFn func(*Autocomplete) error

func (o OptionFn) Configure(rl *Autocomplete) error {
	return o(rl)
}

func New(options ...Option) *Autocomplete {
	rl := &Autocomplete{}
	for _, opt := range options {
		opt.Configure(rl)
	}
	return rl
}

func WithType(typ string) Option {
	return OptionFn(func(rl *Autocomplete) error {
		rl.typ = typ
		return nil
	})
}

func WithValue(typ string) Option {
	return OptionFn(func(rl *Autocomplete) error {
		rl.typ = typ
		return nil
	})
}

func (el *Autocomplete) View() string {

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
				arr = append(arr, key+":='true'")
				continue
			}
			arr = append(arr, key+":='false'")
		}

	}
	return "<el-input v-model='" + el.Getkey() + "' " + strings.Join(arr, " ") + "></el-input>"
}

func (el *Autocomplete) Data() string      { return el.Getkey() + ":\"" + "" + "\"" }
func (el *Autocomplete) Method() string    { return "" }
func (el *Autocomplete) Init() string      { return "" }
func (el *Autocomplete) Computed() string  { return "" }
func (el *Autocomplete) Component() string { return autocompletejs }
