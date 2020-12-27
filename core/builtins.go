package core

import (
	"fmt"
	"math"

	"github.com/adi/hilang-routine/common"
)

// Sqrt ..
var Sqrt = &common.Function{
	Name: "math#sqrt",
	FixedArity: map[int]*common.Overload{
		1: {
			Params: []string{"x"},
			Code: func(env *common.Environment) (interface{}, error) {
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
