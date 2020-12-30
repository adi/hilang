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
	return form.NewFunction(forms)
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
