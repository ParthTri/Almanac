package parser

import (
	"errors"
	"fmt"
	"io"
)

var eof = rune(0)

type Parser struct {
	s   *Scanner
	buf struct {
		tok Token
		lit string
		n   int
	}
}

type TimeRepeat struct {
	Duration float64
	Unit     string
}

type Event struct {
	Date        string
	Name        string
	Time        []string
	TimeRepeat  *TimeRepeat
	Description string
	Tags        []string
}

func NewParser(r io.Reader) *Parser {
	return &Parser{s: NewScanner(r)}
}

// Use the underlying scanner or use the buffer
func (p *Parser) scan() (tok Token, lit string) {
	// If we have a token on the buffer, then return it.
	if p.buf.n != 0 {
		p.buf.n = 0
		return p.buf.tok, p.buf.lit
	}

	// Otherwise read the next token from the scanner.
	tok, lit = p.s.Scan()

	// Save it to the buffer in case we unscan later.
	p.buf.tok, p.buf.lit = tok, lit

	return
}

// unscan pushes the previously read token back onto the buffer.
func (p *Parser) unscan() { p.buf.n = 1 }

func (p *Parser) scanIgnoreWhitespace() (tok Token, lit string) {
	tok, lit = p.scan()
	if tok == WS {
		tok, lit = p.scan()
	}
	return
}

