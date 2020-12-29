package parser

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strconv"
)

// Scanner represents a lexical scanner.
type Scanner struct {
	r *bufio.Reader
}

// NewScanner returns a new instance of Scanner.
func NewScanner(r io.Reader) *Scanner {
	return &Scanner{r: bufio.NewReader(r)}
}

// read reads the next rune from the bufferred reader.
// Returns the rune(0) if an error occurs (or io.EOF is returned).
func (s *Scanner) read() rune {
	ch, _, err := s.r.ReadRune()
	if err != nil {
		return eof
	}
	return ch
}

// unread places the previously read rune back on the reader.
func (s *Scanner) unread() { _ = s.r.UnreadRune() }

// Scan returns the next token and literal value.
func (s *Scanner) Scan() (tok Token, lit interface{}, err error) {
	// Read the next rune.
	ch := s.read()
	// If we see whitespace then consume all contiguous whitespace.
	if isWhitespace(ch) {
		s.unread()
		return s.scanWhitespace()
	}
	// Useful stuff in sight here
	switch ch {
	case eof:
		return EOF, "", nil
	case '(':
		return LPAREN, "", nil
	case ')':
		return RPAREN, "", nil
	case '"':
		s.unread()
		return s.scanString()
	}
	s.unread()
	return s.scanNumberOrSymbol()
}

// scanWhitespace consumes the current rune and all contiguous whitespace.
func (s *Scanner) scanWhitespace() (tok Token, lit interface{}, err error) {
	// Create a buffer and read the current character into it.
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	// Read every subsequent whitespace character into the buffer.
	// Non-whitespace characters and EOF will cause the loop to exit.
	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isWhitespace(ch) {
			s.unread()
			break
		} else {
			buf.WriteRune(ch)
		}
	}

	return WS, buf.String(), nil
}

// scanNumberOrSymbol consumes the current rune and all contiguous number or symbol runes.
func (s *Scanner) scanNumberOrSymbol() (tok Token, lit interface{}, err error) {
	// Create a buffer and read the current character into it.
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	// Read every subsequent symbol character into the buffer.
	// Non-symbol characters and EOF will cause the loop to exit.
	for {
		if ch := s.read(); ch == eof {
			break
		} else if isWhitespace(ch) || isReservedChar(ch) {
			s.unread()
			break
		} else {
			_, _ = buf.WriteRune(ch)
		}
	}

	// Detect if integer, float or symbol
	str := buf.String()

	intVal, err := strconv.ParseInt(str, 10, 64)
	if err == nil {
		return INTEGER, intVal, nil
	}

	floatVal, err := strconv.ParseFloat(str, 64)
	if err == nil {
		return FLOAT, floatVal, nil
	}

	return SYMBOL, buf.String(), nil
}

// scanString consumes the current rune and then until it reaches an unquoted quote.
func (s *Scanner) scanString() (tok Token, lit string, err error) {
	// Create a buffer.
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	// Read every subsequent character into the buffer.
	// Unquoted quote character and EOF will cause the loop to exit.
scanString:
	for {
		ch := s.read()
		switch ch {
		case eof, '"':
			_, _ = buf.WriteRune(ch)
			break scanString
		case '\\':
			escapedCh, err := s.scanStringEscape()
			if err != nil {
				return ILLEGAL, "", err
			}
			_, _ = buf.WriteRune(escapedCh)
		default:
			_, _ = buf.WriteRune(ch)
		}
	}

	return STRING, buf.String(), nil
}

// scanStringEscape consumes until it reaches an unquoted quote.
func (s *Scanner) scanStringEscape() (rune, error) {
	ch := s.read()
	switch ch {
	case 'a':
		return '\a', nil
	case 'b':
		return '\b', nil
	case 'f':
		return '\f', nil
	case 'n':
		return '\n', nil
	case 'r':
		return '\r', nil
	case 't':
		return '\t', nil
	case 'v':
		return '\v', nil
	case '\\':
		return '\\', nil
	case '\'':
		return '\'', nil
	case '"':
		return '"', nil
	case 'u':
		hex := make([]rune, 4)
		for i := 0; i < len(hex); i++ {
			hex[i] = s.read()
		}
		code, err := strconv.ParseInt(string(hex), 16, 16)
		if err != nil {
			return rune(0), fmt.Errorf("couldn't scan 16 bits hex string escape:\n\t%w", err)
		}
		return rune(code), nil
	case 'U':
		hex := make([]rune, 8)
		for i := 0; i < len(hex); i++ {
			hex[i] = s.read()
		}
		code, err := strconv.ParseInt(string(hex), 16, 32)
		if err != nil {
			return rune(0), fmt.Errorf("couldn't scan 32 bits hex string escape:\n\t%w", err)
		}
		return rune(code), nil
	case eof:
		return eof, fmt.Errorf("unexpected EOF within string escape")
	default:
		return ch, fmt.Errorf("unrecognized string escape '\\%c'", ch)
	}
}
