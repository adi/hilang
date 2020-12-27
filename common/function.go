package common

// Function ..
type Function struct {
	Name          string
	FixedArity    map[int]*Overload
	VariadicFixed int
	Variadic      *Overload
}

// Overload ..
type Overload struct {
	Params []string
	Code   func(*Environment) (interface{}, error)
}
