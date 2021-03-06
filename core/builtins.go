package core

import (
	"fmt"
	"math"

	"github.com/adi/hilang-routine/common"
	"github.com/adi/hilang-routine/form"
)

// MathPlus ..
var MathPlus = &form.FunctionForm{
	Name: "+",
	FixedArity: map[int]*form.Overload{
		2: {
			Params: []string{"x", "y"},
			NativeCode: func(env *common.Environment) (interface{}, error) {
				if x, ok := env.Lookup("x"); ok {
					if y, ok := env.Lookup("y"); ok {
						switch x := x.(type) {
						case float64:
							switch y := y.(type) {
							case float64:
								return x + y, nil
							case int64:
								return x + float64(y), nil
							default:
								return nil, fmt.Errorf("+ takes 2 numeric parameters")
							}
						case int64:
							switch y := y.(type) {
							case float64:
								return float64(x) + y, nil
							case int64:
								return x + y, nil
							default:
								return nil, fmt.Errorf("+ takes 2 numeric parameters")
							}
						default:
							return nil, fmt.Errorf("+ takes 2 numeric parameters")
						}
					}
					return nil, fmt.Errorf("internal error: parameter 'y' not found in environment")
				}
				return nil, fmt.Errorf("internal error: parameter 'x' not found in environment")
			},
		},
	},
}

// Sqrt ..
var Sqrt = &form.FunctionForm{
	Name: "math#sqrt",
	FixedArity: map[int]*form.Overload{
		1: {
			Params: []string{"x"},
			NativeCode: func(env *common.Environment) (interface{}, error) {
				if x, ok := env.Lookup("x"); ok {
					switch x := x.(type) {
					case float64:
						return math.Sqrt(x), nil
					case int64:
						return math.Sqrt(float64(x)), nil
					default:
						return nil, fmt.Errorf("sqrt takes a numeric parameter")
					}
				}
				return nil, fmt.Errorf("internal error: parameter 'x' not found in environment")
			},
		},
	},
}
