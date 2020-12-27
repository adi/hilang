package form

// IntegerForm ..
type IntegerForm struct {
	value int64
}

// NewInteger ..
func NewInteger(value int64) *IntegerForm {
	return &IntegerForm{
		value: value,
	}
}

// Eval ..
func (intf IntegerForm) Eval(env *Environment) (interface{}, error) {
	return intf.value, nil
}
