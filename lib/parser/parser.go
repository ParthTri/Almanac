package parser

import (
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

func NewParser(r io.Reader) *Parser {
	return &Parser{s: NewScanner(r)}
}

// func (p *Parser) Parse() error {

// }
