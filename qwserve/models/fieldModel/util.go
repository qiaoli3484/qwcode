package fieldModel

func NewFieldType(typ string) interface{} {
	if typ == "Text" {
		return &Ttext{}
	} else if typ == "Numeric" {
		return &Number{}
	} else if typ == "DateTime" {
		return &Tdate{}
	} else if typ == "SingleList" || typ == "MultipleList" {
		return &SingleList{}
	} else if typ == "YesNo" {
		return &YesNo{}
	} else if typ == "Query" {
		return &Query{}
	} else if typ == "Photo" {
		return &Photo{}
	} else if typ == "File" {
		return &File{}
	} else if typ == "HyperLink" {
		//addLink(c, tid)
	} else if typ == "AutoNumber" {
		return &AutoNumber{}
	} else if typ == "SpecialText" {
		return &SpecialText{}
	} else if typ == "Variable" {
		return &Variable{}
	} else if typ == "Html" {
		return &Html{}
		//addHtml(c, tid)
	}
	return nil
}
