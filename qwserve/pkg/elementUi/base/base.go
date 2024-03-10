package base

type Base struct {
	id   int
	key  string
	name string
}

func (el *Base) SetName(val string) *Base {
	el.name = val
	return el
}
func (el *Base) Setkey(val string) *Base {
	el.key = val
	return el
}
func (el *Base) SetID(val int) { el.id = val }

func (el *Base) GetID(val int)   { el.id = val }
func (el *Base) GetName() string { return el.name }
func (el *Base) Getkey() string  { return el.key }
