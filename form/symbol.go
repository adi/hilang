package form

import (
	"fmt"

	"github.com/adi/hilang-routine/common"
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

// String ..
func (symf SymbolForm) String() string {
	return symf.name
}

// Name ..
func (symf SymbolForm) Name() string {
	return symf.name
}

// Eval ..
func (symf SymbolForm) Eval(env *common.Environment) (interface{}, error) {
	value, ok := env.Lookup(symf.name)
	if ok {
		return value, nil
	}
	return nil, fmt.Errorf("Unable to resolve symbol: %s in this context", symf.name)
}
