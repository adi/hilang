package parser

// Expression can be anything
type Expression interface {
}

// List is a sequence of expressions
type List struct {
	Items []Expression
}

// Integer is the atomic part of code made of only one natural 64 bit number
type Integer struct {
	Literal int64
}

// Float is the atomic part of code made of only one real 64 bit approximated number
type Float struct {
	Literal float64
}

// String is the atomic part of code made of only one string
type String struct {
	Literal string
}

// Symbol is the atomic part of code made of only one name, operator, or anything else
type Symbol struct {
	Literal string
}
