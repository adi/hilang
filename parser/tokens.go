package parser

import "unicode"

// Token is what gets outputed from the lexer
type Token int

// Lexer tokens we might encounter
const (
	ILLEGAL Token = iota // illegal content
	EOF                  // end of code marker
	WS                   // any whitespace
	INTEGER              // natural 64 bit numbers
	FLOAT                // real 64 bit approximations
	STRING               // escaped string style "stuff and \u1047tuff \n"
	SYMBOL               // any unicode except () or STRING, INTEGER or FLOAT
	LPAREN               // (
	RPAREN               // )
)

func isWhitespace(ch rune) bool {
	return unicode.IsSpace(ch)
}

func isReservedChar(ch rune) bool {
	return ch == '(' || ch == ')'
}

var eof = rune(0)
