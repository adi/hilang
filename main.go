package main

import (
	"log"

	"github.com/adi/hilang-routine/common"
	"github.com/adi/hilang-routine/core"
	"github.com/adi/hilang-routine/form"
)

func main() {
	rootEnv := common.NewEnvironment(nil)
	rootEnv.Set("def", core.Def)
	rootEnv.Set("fn", core.Fn)
	rootEnv.Set("do", core.Do)
	rootEnv.Set("+", core.MathPlus)
	rootEnv.Set("sqrt", core.Sqrt)
	program := form.NewList()
	program.Append(form.NewSymbol("do"))

	declaration := form.NewList()
	declaration.Append(form.NewSymbol("def"))
	declaration.Append(form.NewSymbol("mysqrt"))
	fn := form.NewList()
	fn.Append(form.NewSymbol("fn"))
	fnOverload := form.NewList()
	declarationParams := form.NewList()
	declarationParams.Append(form.NewSymbol("x"))
	fnOverload.Append(declarationParams)
	call := form.NewList()
	call.Append(form.NewSymbol("+"))
	call.Append(form.NewSymbol("x"))
	call.Append(form.NewInteger(15))
	fnOverload.Append(call)
	fn.Append(fnOverload)
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
