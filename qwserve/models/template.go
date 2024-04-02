package models

type Template struct {
	ID    int64
	Fid   int64
	Name  string
	Alias string
	Typ   int //1为 流程表 2云盘
	View      //显示
	//Pcview  string `json:"pcview"`  //电脑端视图  云盘地址
	//Phview  string `json:"phview"`  //手机端视图
	//Orderby string `json:"orderby"` //排序字段
	//Sort    string `json:"sort"`    //默认排序
	//Dir     string `json:"dir"`     //显示目录

	Other
	//Unique      bool            `json:"unique"`      //组合唯一字段
	//UnionKeys   []string        `json:"unionKeys"`   //组合唯一字段
	//Envalidity  bool            `json:"envalidity"`  //数据有效性验证
	//Validity    string          `json:"validity"`    //验证提示
	//Validfilter []models.Filter `json:"validfilter"` //验证条件
	//Encomment   bool            `json:"encomment"`   //允许评论
	//Editcomment bool            `json:"editcomment"` //允许修改或删除自己的评论
	//Delcomment  bool            `json:"delcomment"`  //允许记录创建者删除评论

	Protect //保护
	//DisRead   bool            `json:"DisRead"` //禁止浏览
	//DisEdit   bool            `json:"DisEdit"` //禁止编辑
	//DisDel    bool            `json:"DisDel"`  //禁止删除
	//EnRange   bool            `json:"enRange"` //仅保护符合条件的记录
	//Bffilter  []models.Filter `json:"bffilter"`
	//DisAdd    bool            `json:"DisAdd"`    //禁止添加/导入数据
	//DisExport bool            `json:"DisExport"` //禁止导出/备份数据
	//DisPrint  bool            `json:"DisPrint"`  //禁止打印
	//Roles     []int64         `json:"roles"`     //保护例外的角色

	Carrys   []Carry     `json:"carrys"`   //转结内容
	Union    []Union     `json:"union"`    //组合
	Childs   []ChildTemp `json:"childs"`   //子表
	Generate []Generate  `json:"generate"` //用xx生成
	Custom   []Custom    `json:"custom"`   //自定义按钮
	Tbfmt    []Tbfmt
}

// Tbfmt 表格颜色
type Tbfmt struct {
	Flag       string   `json:"flag"`       //来自
	Alias      string   `json:"alias"`      //来自
	Color      string   `json:"color"`      //来自
	Background string   `json:"background"` //来自
	Filter     []Filter `json:"filter"`
}

type Other struct {
	Unique      bool     `json:"unique"`      //组合唯一字段
	UnionKeys   []string `json:"unionKeys"`   //组合唯一字段
	Envalidity  bool     `json:"envalidity"`  //数据有效性验证
	Validity    string   `json:"validity"`    //验证提示
	Validfilter []Filter `json:"validfilter"` //验证条件
	Encomment   bool     `json:"encomment"`   //允许评论
	Editcomment bool     `json:"editcomment"` //允许修改或删除自己的评论
	Delcomment  bool     `json:"delcomment"`  //允许记录创建者删除评论
}

// 显示
type View struct {
	Pcview  string `json:"pcview"`  //电脑端视图  云盘地址
	Phview  string `json:"phview"`  //手机端视图
	Orderby string `json:"orderby"` //排序字段
	Sort    string `json:"sort"`    //默认排序
	//Hide    bool   `json:"hide"`    //显示目录   取消
}

// 保护
type Protect struct {
	DisRead   bool     `json:"DisRead"` //禁止浏览
	DisEdit   bool     `json:"DisEdit"` //禁止编辑
	DisDel    bool     `json:"DisDel"`  //禁止删除
	EnRange   bool     `json:"enRange"` //仅保护符合条件的记录
	Bffilter  []Filter `json:"bffilter"`
	DisAdd    bool     `json:"DisAdd"`    //禁止添加/导入数据
	DisExport bool     `json:"DisExport"` //禁止导出/备份数据
	DisPrint  bool     `json:"DisPrint"`  //禁止打印
	Roles     []int64  `json:"roles"`     //保护例外的角色
}

// Child 子模板
type ChildTemp struct {
	Temp       string   `json:"temp"`       //模板名
	Source     string   `json:"source"`     //对应字段
	Target     string   `json:"target"`     //对应字段
	EnUpdate   bool     `json:"enUpdate"`   //同步更新关联字段
	EnDel      bool     `json:"enDel"`      //同步删除关联记录
	EnAdd      bool     `json:"enAdd"`      //同步添加关联记录 只能添加一条
	EnRequired bool     `json:"enRequired"` //必须填写关联记录
	EnProtect  bool     `json:"enProtect"`  //满足条件时锁定关联记录
	Filter     []Filter `json:"filter"`

	EnTree   bool      `json:"enTree"`   //判断树表
	Tree     int64     `json:"tree"`     //树表Tid
	HasChild string    `json:"hasChild"` //判断是否为树
	Mapping  []Mapping `json:"Mapping"`  //映射
	ParentId string    `json:"parentId"` //对应父ID
}

// Union 组合模板
type Union struct {
	Temp   string   `json:"temp"`
	Source string   `json:"source"`
	Target string   `json:"target"`
	Field  []string `json:"field"`
}

// Mapping 映射
type Mapping struct {
	Source string `json:"source"` //来自
	Typ    string `json:"typ"`    //类型
	Target string `json:"target"` //目标
}

// Carry 自动转结
type Carry struct {
	Id   int64  `json:"id"` //来自结果
	Tid  int64  //来自表
	Name string `json:"name"` //来自结果
	Mode string `json:"mode"` //模式: 统计stats 写入write
	Rtmp int64  `json:"rtmp"`

	Enkey  bool   `json:"enkey"` //是否有主键对应
	Source string `json:"source"`
	Target string `json:"target"`

	RMapping []MappingA `json:"rMapping"`
	Enfilter bool       `json:"enfilter"`
	Moment   string     `json:"moment"`
	Radio    string     `json:"radio"` // 自动转结方式
	Pos      string     `json:"pos"`
	Days     string     `json:"days"`
	Fqcy     string     `json:"fqcy"` //添加 更新  添加或更新
	DelCoTar bool       `json:"delCoTar"`
	DelCoSrc bool       `json:"delCoSrc"`
	Filter   []*Filter  `json:"filter"` //转结条件
	Cron     string
}

// Mapping 映射  //有两个 不一致
type MappingA struct {
	From string `json:"from"` //来自
	Typ  string `json:"typ"`  //类型
	To   string `json:"to"`   //目标
}

type Filter struct {
	Childs []*Child `json:"child"`
}

type Child struct {
	Id       int64  `json:"id"`
	Alias    string `json:"alias"`   //: "Fy77p9l"
	Compare  string `json:"compare"` //: "等于"
	DataType string `json:"dTp"`     //: "AutoNumber"
	Radio    string `json:"radio"`   //: "and"
	Value    string `json:"value"`   //: "KHHK"
	Value2   string `json:"value2"`  //: ""
	Typ      string `json:"typ"`     //custom field
}

type Generate struct {
	Temp    int64     `json:"temp"`
	EnRange bool      `json:"enRange"` //条件
	Filter  []Filter  `json:"filter"`
	Mapping []Mapping `json:"Mapping"` //映射
	Toedit  bool      `json:"Toedit"`  //编辑状态
	Delsrc  bool      `json:"Delsrc"`  //删除源数据
}

// Custom 自定义按钮
type Custom struct {
	Name    string    `json:"name"`   //名称
	Msg     string    `json:"msg"`    //提示
	Type    string    `json:"typ"`    //1置值 2脚本
	Script  string    `json:"script"` //脚本
	Input   bool      `json:"input"`  //输入值脚本
	Roles   []int64   `json:"roles"`  //可使用角色
	EnRange bool      `json:"enRange"`
	Filter  []Filter  `json:"filter"`
	Mapping []Mapping `json:"Mapping"` //映射
}
