package common

// Environment ..
type Environment struct {
	variables map[string]interface{}
	parent    *Environment
}

// NewEnvironment ..
func NewEnvironment(parent *Environment) *Environment {
	return &Environment{
		variables: make(map[string]interface{}),
		parent:    parent,
	}
}

// Set ..
func (he *Environment) Set(name string, val interface{}) {
	he.variables[name] = val
}

// Del ..
func (he *Environment) Del(name string) {
	delete(he.variables, name)
}

// Lookup ..
func (he *Environment) Lookup(name string) (interface{}, bool) {
	if value, ok := he.variables[name]; ok {
		return value, true
	}
	if he.parent != nil {
		return he.parent.Lookup(name)
	}
	return nil, false
}
