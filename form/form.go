package form

import "github.com/adi/hilang-routine/common"

// Form ..
type Form interface {
	Eval(env *common.Environment) (interface{}, error)
}
