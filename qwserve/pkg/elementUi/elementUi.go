package elementUi

import (
	_ "embed"
	"simpleTool/pkg/comm"
	"strings"
)

//go:embed myPlugin.js
var myPlugin string // 当前目录，解析为string类型

type ElementUi struct {
	size   string
	locale string //"zhCn"

	view     strings.Builder
	data     strings.Builder
	methods  strings.Builder
	mounted  strings.Builder
	computed strings.Builder
	use      map[string]string
}

func Default() *ElementUi {
	return &ElementUi{
		use: make(map[string]string),
	}
}

/*func (e *ElementUi) Init(html comm.Htmler) {
	html.AddHead(`<link rel="stylesheet" href="//unpkg.com/element-plus/dist/index.css" />`)
	html.AddHead(`<script src="//unpkg.com/element-plus"></script>`)
}*/

func (e *ElementUi) Config(vue comm.Vuer) {
	vue.AddHead(`<link rel="stylesheet" href="http://unpkg.com/element-plus@2.4.2/dist/index.css" />`)
	vue.AddHead(`<script src="http://unpkg.com/element-plus"></script>`)
	vue.AddHead(`<script src="http://unpkg.com/element-plus/dist/locale/zh-cn"></script>`)
	vue.SetUse("app.use(ElementPlus, {locale: ElementPlusLocaleZhCn,size: 'small', zIndex: 3000 })\n")
	vue.SetUse(myPlugin)
}

func (e *ElementUi) AddEl(el comm.Eler) {
	e.view.WriteString(el.View(""))
	e.data.WriteString(el.Data() + ",")
	e.methods.WriteString(el.Method())
	e.mounted.WriteString(el.Init())
	e.computed.WriteString(el.Computed())

	for kk, vv := range el.Component() {
		e.use[kk] = vv
	}
}

func (e *ElementUi) Component() map[string]string {
	return e.use
}

func (e *ElementUi) Data() string             { return e.data.String() }
func (e *ElementUi) Method() string           { return e.methods.String() }
func (e *ElementUi) Init() string             { return e.mounted.String() }
func (e *ElementUi) Computed() string         { return e.computed.String() }
func (e *ElementUi) View(alias string) string { return e.view.String() }
func (e *ElementUi) GetName() string          { return "elementUi" }
func (e *ElementUi) Getkey() string           { return "elementUi" }
