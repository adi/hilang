package form

import "github.com/adi/hilang-routine/common"

// StringForm ..
type StringForm struct {
	value string
}

// NewString ..
func NewString(value string) *StringForm {
	return &StringForm{
		value: value,
	}
}

// Eval ..
func (strf StringForm) Eval(env *common.Environment) (interface{}, error) {
	return strf.value, nil
}
