package core

import (
	"fmt"

	"github.com/adi/hilang-routine/common"
	"github.com/adi/hilang-routine/form"
)

// Def ..
func Def(env *common.Environment, forms ...form.Form) (interface{}, error) {
	if len(forms) != 2 {
		return nil, fmt.Errorf("def takes exactly 2 parameters")
	}
	variable := forms[0]
	if symbolForm, ok := variable.(*form.SymbolForm); ok {
		symbolName := symbolForm.Name()
		definition := forms[1]
		value, err := definition.Eval(env)
		if err != nil {
			return nil, err
		}
		env.Set(symbolName, value)
		return nil, nil
	}
	return nil, fmt.Errorf("def 1st param must be a symbol")
}

// Fn ..
func Fn(env *common.Environment, forms ...form.Form) (interface{}, error) {
	if len(forms) == 0 {
		return nil, fmt.Errorf("Function has no overloads")
	}
	fnName := ""
	formIndex := 0

	possibleFnName := forms[0]
	if symbolForm, ok := possibleFnName.(*form.SymbolForm); ok {
		if len(forms) == 1 {
			return nil, fmt.Errorf("Function has no overloads")
		}
		fnName = symbolForm.Name()
		formIndex = 1
	}
	retFn := &common.Function{
		Name:       fnName,
		FixedArity: make(map[int]*common.Overload),
	}
	for overloadIndex := formIndex; overloadIndex < len(forms); overloadIndex++ {
		possibleOverload := forms[overloadIndex]
		if overload, ok := possibleOverload.(*form.ListForm); ok {
			overloadItems := overload.Items()
			if len(overloadItems) == 0 {
				return nil, fmt.Errorf("Parameter declaration missing")
			}
			possibleParams := overloadItems[0]
			if params, ok := possibleParams.(*form.ListForm); ok {
				paramsItems := params.Items()
				rawParamNames := make([]string, 0)
				for _, paramForm := range paramsItems {
					if paramSymbolForm, ok := paramForm.(*form.SymbolForm); ok {
						rawParamNames = append(rawParamNames, paramSymbolForm.Name())
					} else {
						return nil, fmt.Errorf("Parameter not a symbol")
					}
				}
				fixedOverload := true
				paramNames := make([]string, 0)
				for i, rawParamName := range rawParamNames {
					if rawParamName == "&" {
						if i != len(rawParamNames)-2 {
							return nil, fmt.Errorf("Variadic separator can only sit at penultimate position in parameter list")
						}
						if retFn.Variadic != nil {
							return nil, fmt.Errorf("Can't have more than 1 variadic overload")
						}
						restParamName := rawParamNames[len(rawParamNames)-1]
						paramNames = append(paramNames, restParamName)
						retFn.Variadic = &common.Overload{
							Params: paramNames,
						}
						retFn.VariadicFixed = i
						fixedOverload = false
						break
					}
					paramNames = append(paramNames, rawParamName)
				}
				if fixedOverload {
					arity := len(paramNames)
					if _, ok := retFn.FixedArity[arity]; ok {
						return nil, fmt.Errorf("Can't have 2 overloads with same arity")
					}
					if retFn.Variadic != nil && arity > retFn.VariadicFixed {
						return nil, fmt.Errorf("Can't have fixed arity function with more params than variadic function")
					}
					retFn.FixedArity[arity] = &common.Overload{
						Params: paramNames,
					}
				}
			} else {
				return nil, fmt.Errorf("Parameter declaration not a list")
			}
		} else {
			return nil, fmt.Errorf("Subform #%d in fn is not an overload", overloadIndex+1)
		}
	}
	return retFn, nil
}

// Do ..
func Do(env *common.Environment, forms ...form.Form) (interface{}, error) {
	var result interface{}
	var err error
	for _, form := range forms {
		result, err = form.Eval(env)
		if err != nil {
			return nil, err
		}
	}
	return result, nil
}
