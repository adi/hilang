package form

import (
	"strings"

	"github.com/adi/hilang-routine/common"
)

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

// String ..
func (strf StringForm) String() string {
	return "\"" + strings.ReplaceAll(strf.value, "\"", "\\\"") + "\""
}

// Eval ..
func (strf StringForm) Eval(env *common.Environment) (interface{}, error) {
	return strf.value, nil
}
