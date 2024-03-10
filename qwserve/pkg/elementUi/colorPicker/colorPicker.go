package colorPicker

import (
	_ "embed"
	"simpleTool/pkg/vue/elementUi/base"
	"simpleTool/pkg/vue/elementUi/util"
)

//var button string // 当前目录，解析为string类型

type elColor struct {
	base.Base

	disabled  bool   `att:"size" default:"false"`      //是否禁用	boolean false
	size      string `att:"size" default:"small"`      //尺寸	enum—
	showAlpha bool   `att:"show-alpha" default:"ture"` //是否支持透明度选择	booleanfalse
	//colorFormat   string `att:"color-format" default:"small"` //写入 v-model 的颜色的格式	enum—
	//predefine     string //预定义颜色	object—
	//validateEvent bool `att:"validate-event" default:"false"` //输入时是否触发表单的校验	booleantrue

	methods string
	data    string
	att     string
}

type Option interface {
	Configure(*elColor) error
}
type OptionFn func(*elColor) error

func (o OptionFn) Configure(rl *elColor) error {
	return o(rl)
}

func New(title string, options ...Option) *elColor {
	rl := &elColor{}
	for _, opt := range options {
		opt.Configure(rl)
	}
	return rl
}

func WithMethods(m string) Option {
	return OptionFn(func(rl *elColor) error {
		//rl.typ = typ
		return nil
	})
}

func (el *elColor) View(alias string) string {
	return "<el-color-picker " + util.AttrRef(el) + el.att + " ></el-button>"
}

func (el *elColor) Data() string {
	/*predefineColors := `predefineColors:[
		'#ff4500',
		'#ff8c00',
		'#ffd700',
		'#90ee90',
		'#00ced1',
		'#1e90ff',
		'#c71585',
		'rgba(255, 69, 0, 0.68)',
		'rgb(255, 120, 0)',
		'hsv(51, 100, 98)',
		'hsva(120, 40, 94, 0.5)',
		'hsl(181, 100%, 37%)',
		'hsla(209, 100%, 56%, 0.73)',
		'#c7158577',
	  ]`*/
	return el.data
}
func (el *elColor) Method() string               { return el.methods }
func (el *elColor) Init() string                 { return "" }
func (el *elColor) Computed() string             { return "" }
func (el *elColor) Component() map[string]string { return nil }
func (el *elColor) GetRule() string              { return "" }
