package form

import (
	"fmt"

	"github.com/adi/hilang-routine/common"
)

// ListForm ..
type ListForm struct {
	items []Form
}

// NewList ..
func NewList() *ListForm {
	return &ListForm{
		items: make([]Form, 0),
	}
}

// Append ..
func (lstf *ListForm) Append(item Form) {
	lstf.items = append(lstf.items, item)
}

// Items ..
func (lstf ListForm) Items() []Form {
	return lstf.items
}

// Eval ..
func (lstf ListForm) Eval(env *common.Environment) (interface{}, error) {
	if len(lstf.items) == 0 {
		return lstf.items, nil
	}
	possibleCallable, err := lstf.items[0].Eval(env)
	if err != nil {
		return nil, fmt.Errorf("Cannot eval first form: %w", err)
	}

	// Special form => pass forms as they are without eval'ing them
	if specialCallable, ok := possibleCallable.(func(*common.Environment, ...Form) (interface{}, error)); ok {
		value, err := specialCallable(env, lstf.items[1:]...)
		if err != nil {
			return nil, fmt.Errorf("Error while invoking: %w", err)
		}
		return value, nil
	}

	// Function => eval args into a function scoped environment before invocation
	if function, ok := possibleCallable.(*common.Function); ok {

		args := lstf.items[1:]

		if fixedOverload, ok := function.FixedArity[len(args)]; ok {
			fnEnv := common.NewEnvironment(env)
			for i, paramName := range fixedOverload.Params {
				value, err := args[i].Eval(env)
				if err != nil {
					return nil, err
				}
				fnEnv.Set(paramName, value)
			}
			return fixedOverload.Code(fnEnv)
		}

		if function.Variadic != nil {
			if len(args) >= function.VariadicFixed {
				fnEnv := common.NewEnvironment(env)
				pLen := len(function.Variadic.Params)
				for i, paramName := range function.Variadic.Params[:pLen-1] {
					value, err := args[i].Eval(env)
					if err != nil {
						return nil, err
					}
					fnEnv.Set(paramName, value)
				}
				fnEnv.Set(function.Variadic.Params[pLen-1], args[:pLen+1])
				return function.Variadic.Code(fnEnv)
			}
		}

		return nil, fmt.Errorf("No overload can take %d arguments", len(args))
	}
	return nil, fmt.Errorf("First form not callable")
}
