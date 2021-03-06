package form

import (
	"strconv"

	"github.com/adi/hilang-routine/common"
)

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

// String ..
func (floatf FloatForm) String() string {
	return strconv.FormatFloat(floatf.value, 'g', -1, 64)
}

// Eval ..
func (floatf FloatForm) Eval(env *common.Environment) (interface{}, error) {
	return floatf.value, nil
}
