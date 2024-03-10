package table

import (
	"fmt"
	"simpleTool/pkg/comm"
	"simpleTool/pkg/vue/elementUi/base"
	"simpleTool/pkg/vue/elementUi/util"
	"strings"
)

type ElTable struct {
	base.Base
	size                string `att:"size" default:"small"`
	border              bool   `att:"border" default:"true"`
	width               string `att:"width" default:"100%%"`
	height              string `att:"height" default:"auto"`
	highlightCurrentRow string `att:"highlight-current-row" default:"true"`

	view     strings.Builder
	data     strings.Builder
	methods  strings.Builder
	mounted  strings.Builder
	computed strings.Builder
	use      strings.Builder

	rows []*Row
}

type Row struct {
	el                  comm.Eler
	typ                 string //对应列的类型。 如果设置了selection则显示多选框； 如果设置了 index 则显示该行的索引（从 1 开始计算）； 如果设置了 expand 则显示为一个可展开的按钮	string	selection / index / expand	—
	index               int    //如果设置了 type=index，可以通过传递 index 属性来自定义索引	number / function(index)	—	—
	label               string //显示的标题		—	—
	columnKey           string //column 的 key， column 的 key， 如果需要使用 filter-change 事件，则需要此属性标识是哪个 column 的筛选条件		—	—
	prop                string //字段名称 对应列内容的字段名， 也可以使用 property属性		—	—
	width               string //对应列的宽度	 / number	—	—
	minWidth            string //对应列的最小宽度， 对应列的最小宽度， 与 width 的区别是 width 是固定的，min-width 会把剩余宽度按比例分配给设置了 min-width 的列	 / number	—	—
	fixed               string //列是否固定在左侧或者右侧。 true 表示固定在左侧	string / boolean	true / 'left' / 'right'	—
	renderHeader        string //列标题 Label 区域渲染使用的 Function	function({ column, $index })	—	—
	sortable            string //对应列是否可以排序， 如果设置为 'custom'，则代表用户希望远程排序，需要监听 Table 的 sort-change 事件	boolean / string	custom	false
	sortMethod          string //指定数据按照哪个属性进行排序，仅当sortable设置为true的时候有效。 应该如同 Array.sort 那样返回一个 Number	function(a, b)	—	—
	sortBy              string //指定数据按照哪个属性进行排序，仅当 sortable 设置为 true 且没有设置 sort-method 的时候有效。 如果 sort-by 为数组，则先按照第 1 个属性排序，如果第 1 个相等，再按照第 2 个排序，以此类推	function(row, index) / string / array	—	—
	sortOrders          string //数据在排序时所使用排序策略的轮转顺序，仅当 sortable 为 true 时有效。 需传入一个数组，随着用户点击表头，该列依次按照数组中元素的顺序进行排序	array	数组中的元素需为以下三者之一：ascending 表示升序，descending 表示降序，null 表示还原为原始顺序	['ascending', 'descending', null]
	resizable           bool   //对应列是否可以通过拖动改变宽度（需要在 el-table 上设置 border 属性为真）	boolean	—	true
	formatter           string //用来格式化内容	function(row, column, cellValue, index)	—	—
	showOverflowTooltip bool   //当内容过长被隐藏时显示 tooltip	boolean \
	align               string //对齐方式	string	left / center / right	left
	headerAlign         string //表头对齐方式， 若不设置该项，则使用表格的对齐方式	string	left / center / right	—
	className           string //列的 className	string	—	—
	labelClassName      string //当前列标题的自定义类名	string	—	—
	selectable          string //仅对 type=selection 的列有效，类型为 Function，Function 的返回值用来决定这一行的 CheckBox 是否可以勾选	function(row, index)	—	—
	reserveSelection    bool   //数据刷新后是否保留选项，仅对  type=selection 的列有效， 请注意， 需指定 row-key 来让这个功能生效。	boolean	—	false
	filters             string //数据过滤的选项， 数组格式，数组中的元素需要有 text 和 value 属性。 数组中的每个元素都需要有 text 和 value 属性。	Array<{text: string, value: string}>	—	—
	filterPlacement     string //过滤弹出框的定位	string	与 Tooltip 的 placement 属性相同	—
	filterMultiple      bool   //数据过滤的选项是否多选	boolean	—	true
	filterMethod        string //数据过滤使用的方法， 如果是多选的筛选项，对每一条数据会执行多次，任意一次返回 true 就会显示。	function(value, row, column)	—	—
	filteredValue       string //选中的数据过滤项，如果需要自定义表头过滤的渲染方式，可能会需要此属性。	array
}

type Option interface {
	Configure(*ElTable) error
}
type OptionFn func(*ElTable) error

func (o OptionFn) Configure(rl *ElTable) error {
	return o(rl)
}

func New(key string, options ...Option) *ElTable {
	rl := &ElTable{}
	rl.Setkey(key)
	for _, opt := range options {
		opt.Configure(rl)
	}
	return rl
}

func (b *ElTable) Default() *ElTable {
	b.methods.WriteString(fmt.Sprintf(`
	submitForm%[1]s() {
        this.$refs['ref_%[1]s'].validate((valid) => {
          if (valid) {
            alert('submit!');
          } else {
            console.log('error submit!!');
            return false;
          }
        });
      },
      resetForm%[1]s(formName) {
        this.$refs['ref_%[1]s'].resetFields();
      }`, b.Getkey()))

	return b
}

func (b *ElTable) Index() *ElTable {
	b.view.WriteString(`<el-table-column type="index" width="50" />`)
	return b
}

func (b *ElTable) Multiple() *ElTable {
	b.view.WriteString(`<el-table-column type="selection" width="55" />`)
	return b
}

func (t *ElTable) AddEl(el comm.Eler, width string, IsRender bool) {
	if IsRender {
		t.view.WriteString("<el-table-column prop=\"" + el.Getkey() + "\" label=\"" + el.GetName() + "\" align=\"center\" sortable show-overflow-tooltip width=\"" + width + "\" /><template #default=\"scope\">")
		t.view.WriteString(el.View("scope.row"))
		//<template #default="scope">{{ scope.row.date }}</template>
		t.view.WriteString("</template></el-table-column>")
	} else {
		t.view.WriteString("<el-table-column prop=\"" + el.Getkey() + "\" label=\"" + el.GetName() + "\" align=\"center\" sortable show-overflow-tooltip width=\"" + width + "\" />")
	}
	t.data.WriteString("," + el.Data())
}

func (el *ElTable) View(alias string) string {
	if alias != "" {
		alias = alias + "."
	}

	head := fmt.Sprintf(`<el-table ref="ref_%[2]s" :data="%[1]s" `+util.AttrRef(el)+`>`, alias+el.Getkey(), el.Getkey())

	return head + el.view.String() + "</el-table>"
}

func (el *ElTable) Data() string {
	head := el.Getkey() + ":[{ID:0"

	return head + el.data.String() + "}]"
}

func (el *ElTable) Method() string   { return "" }
func (el *ElTable) Init() string     { return "" }
func (el *ElTable) Computed() string { return "" }
func (el *ElTable) Component() map[string]string {
	return nil
}

func (el *ElTable) GetRule() string {
	return ""
}
