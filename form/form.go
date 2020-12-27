package form

// Form ..
type Form interface {
	Eval(env *Environment) (interface{}, error)
}
