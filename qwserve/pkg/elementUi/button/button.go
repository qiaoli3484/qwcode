package button

import (
	_ "embed"
	"fmt"
	"qwserve/pkg/elementUi/base"
	"qwserve/pkg/elementUi/util"
	"simpleTool/pkg/comm"
	"strings"
)

//var button string // 当前目录，解析为string类型

type elInput struct {
	base.Base

	title       string
	size        string `att:"size" default:"small"`
	typ         string `att:"type" default:"primary"`
	plain       bool
	round       bool
	circle      bool
	loading     bool
	disabled    bool
	icon        string
	autofocus   bool
	native_type string

	methods string
	data    string
	att     string
}

type Option interface {
	Configure(*elInput) error
}
type OptionFn func(*elInput) error

func (o OptionFn) Configure(rl *elInput) error {
	return o(rl)
}

func New(title string, options ...Option) *elInput {
	rl := &elInput{
		title: title,
	}
	rl.typ = "primary"
	for _, opt := range options {
		opt.Configure(rl)
	}
	return rl
}

func WithType(typ string) Option {
	return OptionFn(func(rl *elInput) error {
		rl.typ = typ
		return nil
	})
}

func WithMethods(m string) Option {
	return OptionFn(func(rl *elInput) error {
		if m == "save" {
			rl.att = "@click=\"save()\" :loading='saveload'"
			rl.methods = `async save(){
			 this.saveload=true
			 let valid=false
			 valid=await this.submitForm().catch(()=>{
				console.log('error submit!!');
				this.saveload=false
				return
			 })
			 if (valid) {
				alert('submit!');
			  } else {
				console.log('error submit!!');
				return false;
			  }
			 this.saveload=false
			 console.log("哈哈哈")
		},`
		} else if m == "new" {
			rl.att = "@click=\"resetForm()\" :loading='saveload'"
			//id=0 ??
		}

		rl.data = "saveload:false"
		//rl.typ = typ
		return nil
	})
}

func (el *elInput) View(alias string) string {

	return fmt.Sprintf("<el-button "+util.AttrRef(el)+el.att+" >%s</el-button>", el.title)
}

func (el *elInput) Data() string                 { return el.data }
func (el *elInput) Method() string               { return el.methods }
func (el *elInput) Init() string                 { return "" }
func (el *elInput) Computed() string             { return "" }
func (el *elInput) Component() map[string]string { return nil }
func (el *elInput) GetRule() string              { return "" }

// <el-button-group>
type elGroup struct {
	base.Base

	view    strings.Builder
	data    strings.Builder
	methods strings.Builder
}

func Group(els ...comm.FEler) *elGroup {
	g := &elGroup{}
	for _, el := range els {
		g.view.WriteString(el.View(""))
		g.methods.WriteString(el.Method())
	}

	return g
}

func (el *elGroup) View(alias string) string {
	return "<el-button-group>" + el.view.String() + "</el-button-group>"
}

func (el *elGroup) Data() string                 { return "" }
func (el *elGroup) Method() string               { return el.methods.String() }
func (el *elGroup) Init() string                 { return "" }
func (el *elGroup) Computed() string             { return "" }
func (el *elGroup) Component() map[string]string { return nil }
func (el *elGroup) GetRule() string              { return "" }
