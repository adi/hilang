package form

import "github.com/adi/hilang-routine/common"

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
func (intf IntegerForm) Eval(env *common.Environment) (interface{}, error) {
	return intf.value, nil
}
