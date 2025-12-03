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

// func (p *Parser) Parse() error {

// }
