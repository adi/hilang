package main

import (
	"fmt"
	"log"
	"math"

	"github.com/adi/hilang-routine/form"
)

func def(env *form.Environment, forms ...form.Form) (interface{}, error) {
	if len(forms) != 2 {
		return nil, fmt.Errorf("def takes exactly 2 arguments")
	}
	variable := forms[0]
	if symbolForm, ok := variable.(*form.SymbolForm); ok {
		symbolName := symbolForm.Name()
		definition := forms[1]
		env.Set(symbolName, definition)
		return nil, nil
	}
	return nil, fmt.Errorf("def 1st arg must be a symbol")
}

func do(env *form.Environment, forms ...form.Form) (interface{}, error) {
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

func sqrt(env *form.Environment, values ...interface{}) (interface{}, error) {
	if len(values) != 1 {
		return nil, fmt.Errorf("sqrt takes exactly one argument")
	}
	value := values[0]
	var err error
	if expr, ok := value.(form.Form); ok {
		value, err = expr.Eval(env)
		if err != nil {
			return nil, fmt.Errorf("Can't evaluate argument")
		}
	}
	switch value := value.(type) {
	case float64:
		return math.Sqrt(value), nil
	case int64:
		return math.Sqrt(float64(value)), nil
	default:
		return nil, fmt.Errorf("sqrt takes a numeric argument")
	}
}

func main() {
	rootEnv := form.NewEnvironment(nil)
	rootEnv.Set("do", do)
	rootEnv.Set("def", def)
	rootEnv.Set("sqrt", sqrt)
	program := form.NewList()
	program.Append(form.NewSymbol("do"))
	declaration := form.NewList()
	declaration.Append(form.NewSymbol("def"))
	declaration.Append(form.NewSymbol("mysqrt"))
	fn := form.NewList()
	fn.Append(form.NewSymbol("fn"))
	declarationParams := form.NewList()
	declarationParams.Append(form.NewSymbol("x"))
	fn.Append(declarationParams)
	call := form.NewList()
	call.Append(form.NewSymbol("sqrt"))
	call.Append(form.NewSymbol("x"))
	fn.Append(call)
	declaration.Append(fn)
	program.Append(declaration)
	statement := form.NewList()
	statement.Append(form.NewSymbol("mysqrt"))
	statement.Append(form.NewInteger(2))
	program.Append(statement)
	val, err := program.Eval(rootEnv)
	if err != nil {
		panic(err)
	}
	log.Printf("Result is: %v\n", val)
}
