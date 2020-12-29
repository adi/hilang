package form

import "github.com/adi/hilang-routine/common"

// Function ..
type Function struct {
	Name          string
	FixedArity    map[int]*Overload
	VariadicFixed int
	Variadic      *Overload
}

// Overload ..
type Overload struct {
	Params     []string
	Code       []Form
	NativeCode func(env *common.Environment) (interface{}, error)
}
