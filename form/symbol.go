package form

import (
	"fmt"
)

// SymbolForm ..
type SymbolForm struct {
	name string
}

// NewSymbol ..
func NewSymbol(name string) *SymbolForm {
	return &SymbolForm{
		name: name,
	}
}

// Name ..
func (s SymbolForm) Name() string {
	return s.name
}

// Eval ..
func (s SymbolForm) Eval(env *Environment) (interface{}, error) {
	value, ok := env.Lookup(s.name)
	if ok {
		return value, nil
	}
	return nil, fmt.Errorf("Unable to resolve symbol: %s in this context", s.name)
}
