package form

import (
	"fmt"
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

// Eval ..
func (lstf ListForm) Eval(env *Environment) (interface{}, error) {
	if len(lstf.items) == 0 {
		return lstf.items, nil
	}
	possibleCallable, err := lstf.items[0].Eval(env)
	if err != nil {
		return nil, fmt.Errorf("Cannot eval first form: %w", err)
	}
	// Special form => pass forms as they are without eval'ing them
	if specialCallable, ok := possibleCallable.(func(*Environment, ...Form) (interface{}, error)); ok {
		value, err := specialCallable(env, lstf.items[1:]...)
		if err != nil {
			return nil, fmt.Errorf("Error while invoking: %w", err)
		}
		return value, nil
	}
	// Builtin function form => eval args before invocation
	if builtinCallable, ok := possibleCallable.(func(*Environment, ...interface{}) (interface{}, error)); ok {
		args := make([]interface{}, 0)
		for _, item := range lstf.items[1:] {
			arg, err := item.Eval(env)
			if err != nil {
				return nil, fmt.Errorf("Cannot eval arg: %w", err)
			}
			args = append(args, arg)
		}
		value, err := builtinCallable(env, args...)
		if err != nil {
			return nil, fmt.Errorf("Error while invoking: %w", err)
		}
		return value, nil
	}
	// Form => verify it can be called
	if possibleCallableForm, ok := possibleCallable.(Form); ok {
		if listForm, ok := possibleCallableForm.(*ListForm); ok {
			if len(listForm.items) > 0 {
				if symbolForm, ok := listForm.items[0].(*SymbolForm); ok {
					if symbolForm.name == "fn" {
						if len(listForm.items) < 3 {
							return nil, fmt.Errorf("fn takes at least 2 parameters")
						}
						if paramListForm, ok := listForm.items[1].(*ListForm); ok {
							allSymbols := true
							for _, paramForm := range paramListForm.items {
								if _, ok := paramForm.(*SymbolForm); !ok {
									allSymbols = false
									break
								}
							}
							if !allSymbols {
								return nil, fmt.Errorf("fn takes a list of symbols as 1st parameter")
							}
							fnEnv := NewEnvironment(env)
							args := lstf.items[1:]
							for i, paramForm := range paramListForm.items {
								if paramSymbolForm, ok := paramForm.(*SymbolForm); ok {
									fnEnv.Set(paramSymbolForm.name, args[i])
								}
							}
							stmtListForm := listForm.items[2:]
							var result interface{}
							var err error
							for _, stmtForm := range stmtListForm {
								result, err = stmtForm.Eval(fnEnv)
								if err != nil {
									return nil, err
								}
							}
							return result, nil
						}
						return nil, fmt.Errorf("fn takes a list as 1st argument")
					}
				}
			}
		}
	}
	return nil, fmt.Errorf("First form not callable")
}
