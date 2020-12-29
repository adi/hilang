package parser

import (
	"fmt"
	"io"
)

// Parser represents a parser.
type Parser struct {
	s        *Scanner
	filename string
	buf      struct {
		tok Token       // last read token
		lit interface{} // last read literal
		n   int         // buffer size (max=1)
	}
}

// NewParser returns a new instance of Parser.
func NewParser(r io.Reader, filename string) *Parser {
	return &Parser{s: NewScanner(r), filename: filename}
}

// scan returns the next token from the underlying scanner.
// If a token has been unscanned then read that instead.
func (p *Parser) scan() (tok Token, lit interface{}, err error) {
	// If we have a token on the buffer, then return it.
	if p.buf.n != 0 {
		p.buf.n = 0
		return p.buf.tok, p.buf.lit, nil
	}

	// Otherwise read the next token from the scanner.
	for {
		tok, lit, err = p.s.Scan()
		if err != nil {
			return ILLEGAL, "", fmt.Errorf("couldn't scan token:\n\t%w", err)
		}
		if tok != WS {
			break
		}
	}

	// Save it to the buffer in case we unscan later.
	p.buf.tok, p.buf.lit = tok, lit

	return
}

// unscan pushes the previously read token back onto the buffer.
func (p *Parser) unscan() { p.buf.n = 1 }

// Parse is transforming the contents from the contained io.Reader into an AST
func (p *Parser) Parse() ([]Expression, error) {
	return p.parseExpressionList(false)
}

func (p *Parser) parseExpressionList(subExpression bool) ([]Expression, error) {
	items := make([]Expression, 0)
parseExpression:
	for {
		tok, lit, err := p.scan()
		if err != nil {
			return nil, fmt.Errorf("couldn't scan in expression list:\n\t%w", err)
		}
		switch tok {
		case EOF:
			if subExpression {
				return nil, fmt.Errorf("unexpected EOF in subexpression instead of item or )")
			}
			break parseExpression
		case RPAREN:
			if subExpression {
				break parseExpression
			}
			return nil, fmt.Errorf("unexpected ) in expression list")
		case INTEGER:
			items = append(items, &Integer{Literal: lit.(int64)})
		case FLOAT:
			items = append(items, &Float{Literal: lit.(float64)})
		case SYMBOL:
			items = append(items, &Symbol{Literal: lit.(string)})
		case LPAREN:
			subexpressionList, err := p.parseExpressionList(true)
			if err != nil {
				return nil, fmt.Errorf("couldn't parse subexpression:\n\t%w", err)
			}
			items = append(items, subexpressionList)
		}
	}
	return items, nil
}
