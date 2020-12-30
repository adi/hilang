package form

import (
	"fmt"

	"github.com/adi/hilang-routine/common"
)

// FunctionForm ..
type FunctionForm struct {
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

// NewFunction ..
func NewFunction(forms []Form) (*FunctionForm, error) {
	if len(forms) == 0 {
		return nil, fmt.Errorf("Function has no overloads")
	}
	fnName := ""
	formIndex := 0

	possibleFnName := forms[0]
	if symbolForm, ok := possibleFnName.(*SymbolForm); ok {
		if len(forms) == 1 {
			return nil, fmt.Errorf("Function has no overloads")
		}
		fnName = symbolForm.Name()
		formIndex = 1
	}
	retFn := &FunctionForm{
		Name:       fnName,
		FixedArity: make(map[int]*Overload),
	}
	for overloadIndex := formIndex; overloadIndex < len(forms); overloadIndex++ {
		possibleOverload := forms[overloadIndex]
		if confirmedOverload, ok := possibleOverload.(*ListForm); ok {
			overloadItems := confirmedOverload.Items()
			if len(overloadItems) == 0 {
				return nil, fmt.Errorf("Parameter declaration missing")
			}
			possibleParams := overloadItems[0]
			bodyForms := overloadItems[1:]
			if params, ok := possibleParams.(*ListForm); ok {
				paramsItems := params.Items()
				rawParamNames := make([]string, 0)
				for _, paramForm := range paramsItems {
					if paramSymbolForm, ok := paramForm.(*SymbolForm); ok {
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
						retFn.Variadic = &Overload{
							Params: paramNames,
							Code:   bodyForms,
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
					retFn.FixedArity[arity] = &Overload{
						Params: paramNames,
						Code:   bodyForms,
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

// String ..
func (fn FunctionForm) String() string {
	ret := "(fn "
	if fn.Name != "" {
		ret += fn.Name + " "
	}
	for _, overload := range fn.FixedArity {
		ret += "( "
		ret += "( "
		for _, param := range overload.Params {
			ret += param + " "
		}
		ret += ")"
		if overload.Code != nil {
			for _, form := range overload.Code {
				ret += form.String() + " "
			}
		} else {
			ret += "(core#native "
			ret += fmt.Sprintf("%p", &overload.NativeCode)
			ret += ")"
		}
		ret += ")"
	}
	if fn.Variadic != nil {
		ret += "( "
		ret += "( "
		for i, param := range fn.Variadic.Params {
			if i == len(fn.Variadic.Params)-1 {
				ret += "& "
			}
			ret += param + " "
		}
		ret += ")"
		if fn.Variadic.Code != nil {
			for _, form := range fn.Variadic.Code {
				ret += form.String() + " "
			}
		} else {
			ret += "(core#native "
			ret += fmt.Sprintf("%p", &fn.Variadic.NativeCode)
			ret += ")"
		}
		ret += ")"
	}
	ret += ")"
	return ret
}

// Eval ..
func (fn FunctionForm) Eval(env *common.Environment) (interface{}, error) {
	return fn, nil
}
