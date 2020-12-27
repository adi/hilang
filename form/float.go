package form

import "github.com/adi/hilang-routine/common"

// FloatForm ..
type FloatForm struct {
	value float64
}

// NewFloat ..
func NewFloat(value float64) *FloatForm {
	return &FloatForm{
		value: value,
	}
}

// Eval ..
func (floatf FloatForm) Eval(env *common.Environment) (interface{}, error) {
	return floatf.value, nil
}
