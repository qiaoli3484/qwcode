package qlselect

import (
	_ "embed"
	"fmt"
	"simpleTool/pkg/vue/elementUi/base"
	"simpleTool/pkg/vue/elementUi/util"
	"strings"
)

//go:embed qlselect.js
var qlselect string // 当前目录，解析为string类型

type Select struct {
	base.Base
	size     string `att:"size" default:"small"`
	multiple bool   `att:"multiple" default:"false"`

	options []string `att:"options"`
	vSHOW   string

	data strings.Builder
}

type Option interface {
	Configure(*Select) error
}

type OptionFn func(*Select) error

func (o OptionFn) Configure(rl *Select) error {
	return o(rl)
}

func New(data string, options ...Option) *Select {
	rl := &Select{}
	for _, opt := range options {
		opt.Configure(rl)
	}
	arr := strings.Split(data, "|")
	rl.options = arr
	return rl
}

func WithValue(typ string) Option {
	return OptionFn(func(rl *Select) error {
		//rl.typ = typ
		return nil
	})
}

func (el *Select) SetName(val string) *Select {
	el.Base.SetName(val)
	return el
}

func (el *Select) Setkey(val string) *Select {
	el.Base.Setkey(val)
	return el
}

func (el *Select) SetShow(val string) *Select {
	el.vSHOW = val
	return el
}

func (el *Select) View(alias string) string {

	if alias != "" {
		alias = alias + "."
	}

	return "<ql-select v-model:data='" + alias + el.Getkey() + "' " + util.AttrRef(el) + ">" + el.data.String() + "</ql-select>"
}

func (el *Select) Data() string     { return el.Getkey() + ":\"" + "" + "\"" }
func (el *Select) Method() string   { return "" }
func (el *Select) Init() string     { return "" }
func (el *Select) Computed() string { return "" }
func (el *Select) Component() map[string]string {
	return map[string]string{"selectjs": qlselect}
}

func (el *Select) GetRule() string {
	s1 := `%s: [
	{ required: true, message: '请选择%s', trigger: 'change' }
  ],`
	return fmt.Sprintf(s1, el.Getkey(), el.GetName())
}
