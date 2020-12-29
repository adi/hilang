package main

import (
	"log"
	"strings"

	"github.com/adi/hilang-routine/common"
	"github.com/adi/hilang-routine/core"
	"github.com/adi/hilang-routine/form"
	"github.com/adi/hilang-routine/parser"
)

func transformExpressionIntoForm(expr parser.Expression) form.Form {
	switch expr := expr.(type) {
	case *parser.Integer:
		return form.NewInteger(expr.Literal)
	case *parser.Float:
		return form.NewFloat(expr.Literal)
	case *parser.String:
		return form.NewString(expr.Literal)
	case *parser.Symbol:
		return form.NewSymbol(expr.Literal)
	case []parser.Expression:
		exprAsList := form.NewList()
		for _, item := range expr {
			exprAsList.Append(transformExpressionIntoForm(item))
		}
		return exprAsList
	}
	return nil
}

func main() {
	rootEnv := common.NewEnvironment(nil)
	rootEnv.Set("def", core.Def)
	rootEnv.Set("fn", core.Fn)
	rootEnv.Set("do", core.Do)
	rootEnv.Set("+", core.MathPlus)
	rootEnv.Set("sqrt", core.Sqrt)
	program := form.NewList()
	program.Append(form.NewSymbol("do"))

	code := `
		(def mysqrt
			(fn ((x) (+ x 15))))
		(mysqrt 2)
	`

	parser := parser.NewParser(strings.NewReader(code), "main.hi")
	expressions, err := parser.Parse()
	if err != nil {
		panic(err)
	}

	for _, expr := range expressions {
		program.Append(transformExpressionIntoForm(expr))
	}

	val, err := program.Eval(rootEnv)

	if err != nil {
		panic(err)
	}
	log.Printf("Result is: %v\n", val)
}
