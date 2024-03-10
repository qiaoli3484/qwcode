package checkbox

import (
	_ "embed"
	"fmt"
	"reflect"
	"simpleTool/pkg/vue/elementUi/base"
	"strconv"
	"strings"
)

//go:embed checkbox.js
var checkboxjs string // 当前目录，解析为string类型

type Checkbox struct {
	base.Base
	size     string   `att:"size" default:"small"`
	disabled bool     `att:"disabled"`
	options  []string `att:"options"`
	vSHOW    string

	data strings.Builder
}

type Option interface {
	Configure(*Checkbox) error
}
type OptionFn func(*Checkbox) error

func (o OptionFn) Configure(rl *Checkbox) error {
	return o(rl)
}

func New(data string, options ...Option) *Checkbox {
	rl := &Checkbox{}
	for _, opt := range options {
		opt.Configure(rl)
	}
	arr := strings.Split(data, "|")
	rl.options = arr
	return rl
}

func WithValue(typ string) Option {
	return OptionFn(func(rl *Checkbox) error {
		//rl.typ = typ
		return nil
	})
}

func (el *Checkbox) SetName(val string) *Checkbox {
	el.Base.SetName(val)
	return el
}

func (el *Checkbox) Setkey(val string) *Checkbox {
	el.Base.Setkey(val)
	return el
}

func (el *Checkbox) SetShow(val string) *Checkbox {
	el.vSHOW = val
	return el
}

func (el *Checkbox) View(alias string) string {

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
				arr = append(arr, key+"=\""+def+"\"")
				continue
			}
			arr = append(arr, key+"=\""+sv.String()+"\"")

		} else if sf.Type.Kind() == reflect.Int {
			if sv.Int() == 0 {
				arr = append(arr, ":"+key+"=\""+def+"\"")
				continue
			}
			arr = append(arr, ":"+key+"=\""+strconv.Itoa(int(sv.Int()))+"\"")
		} else if sf.Type.Kind() == reflect.Bool {
			bb, _ := strconv.ParseBool(def)

			if sv.Bool() || bb {
				arr = append(arr, key)
				continue
			}
			arr = append(arr, ":"+key+"='false'")
		} else if sf.Type.Kind() == reflect.Slice {

			if sv.Len() == 0 {
				arr = append(arr, ":"+key+"=\"[]\"")
				continue
			}
			ar1 := make([]string, sv.Len())
			for i := 0; i < sv.Len(); i++ {
				v := sv.Index(i)
				ar1[i] = v.String()
			}
			arr = append(arr, ":"+key+"=\"['"+strings.Join(ar1, " ','")+"']\"")
			//sv.Elem().Slice()
		}

		//"['" + strings.Join(arr, "','") + "']"
	}

	return "<ql-checkbox v-model:data='" + alias + el.Getkey() + "' " + strings.Join(arr, " ") + ">" + el.data.String() + "</ql-checkbox>"
}

func (el *Checkbox) Data() string     { return el.Getkey() + ":\"" + "" + "\"" }
func (el *Checkbox) Method() string   { return "" }
func (el *Checkbox) Init() string     { return "" }
func (el *Checkbox) Computed() string { return "" }
func (el *Checkbox) Component() map[string]string {
	return map[string]string{"checkboxjs": checkboxjs}
}

func (el *Checkbox) GetRule() string {
	s1 := `%s: [
	{ required: true, message: '请选择%s', trigger: 'change' }
  ],`
	return fmt.Sprintf(s1, el.Getkey(), el.GetName())
}
