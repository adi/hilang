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

// Let ..
func Let(env *common.Environment, forms ...form.Form) (interface{}, error) {
	if len(forms) < 1 {
		return nil, fmt.Errorf("let takes at least 1 parameter")
	}
	probableBindings := forms[0]
	if bindings, ok := probableBindings.(*form.ListForm); ok {
		bindingItems := bindings.Items()
		if len(bindingItems)%2 != 0 {
			return nil, fmt.Errorf("let requires an even number of forms in binding list")
		}
		letEnv := common.NewEnvironment(env)
		for b := 0; b < len(bindingItems)/2; b++ {
			key := ""
			keyForm := bindingItems[2*b]
			if keyForm, ok := keyForm.(*form.SymbolForm); ok {
				key = keyForm.Name()
			} else {
				return nil, fmt.Errorf("each pair in binding list has to start with a symbol")
			}
			value, err := bindingItems[2*b+1].Eval(letEnv)
			if err != nil {
				return nil, err
			}
			letEnv.Set(key, value)
		}
		var result interface{}
		var err error
		for _, form := range forms[1:] {
			result, err = form.Eval(letEnv)
			if err != nil {
				return nil, err
			}
		}
		return result, nil
	}
	return nil, fmt.Errorf("let 1st param must be a list")
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
