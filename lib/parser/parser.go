package parser

import (
	"errors"
	"fmt"
	"io"
	"slices"
	"strings"
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

func (ev *Event) setName(p *Parser, tok Token, lit string) {
	name := lit
	for {
		tok, lit = p.scan()
		if tok == WS {
			name += lit
		} else if tok == WORD {
			name += lit
		} else {
			p.unscan()
			break
		}
	}
	ev.Name = strings.TrimSpace(name)
}

func (p *Parser) parseSingle(ev *Event) error {
	event := ev

	// Scan the time
	tok, lit := p.scanIgnoreWhitespace()
	if tok == EOL {
		tok, lit = p.scanIgnoreWhitespace()
	}

	// Hits 2 consecutive EOL then break
	if tok == EOL || tok == EOF {
		return errors.New("EOF")
	} else if tok != TIME {
		return errors.New(fmt.Sprintf("Expected TIME got %v", tok))
	} else {
		event.Time = append(event.Time, lit)
	}

	// Check if it is a time range
	tok, lit = p.scanIgnoreWhitespace()
	if tok == DASH {
		tok, lit = p.scanIgnoreWhitespace()
		if tok != TIME {
			return errors.New(fmt.Sprintf("Expected TIME for range got %v", lit))
		}

		event.Time = append(event.Time, lit)
	}

	// Check for name
	if tok != WORD {
		tok, lit = p.scanIgnoreWhitespace()
	}
	event.setName(p, tok, lit)

	for {
		// Check for tag
		// OR
		// Check for description
		tok, lit = p.scanIgnoreWhitespace()
		if tok == EOL {
			break
		} else if tok == TAG {
			tok, lit = p.scanIgnoreWhitespace()
			event.Tags = append(event.Tags, lit)
		} else if tok == DESCRIPTION {
			description := ""
			_, lit := p.scanIgnoreWhitespace()
			for tok != EOL {
				description += lit
				tok, lit = p.scan()
			}
			p.unscan()

			event.Description = description
		}
	}

	return nil
}

func (p *Parser) Parse() ([]*Event, error) {
	/*
		Parse using the given scanner to return the first Event struct filled in

		Name should be sequence of words seperated by whitespace before a TAG, DESCRIPTION or EOL
	*/

	// TODO: Check if the input is a multi-line entry or a single line entry and run according to that
	events := []*Event{}
	event := &Event{}
	var err error

	tok1, lit1 := p.scan()
	tok2, _ := p.scan()
	if tok1 == EOF {
		return []*Event{nil}, nil
	}

	if tok1 == DATE && tok2 == EOL {
		event.Date = lit1
		err = p.parseSingle(event)

		for err == nil {
			events = append(events, event)
			event = &Event{}
			event.Date = lit1
			err = p.parseSingle(event)
		}

	} else if tok1 == DATE && tok2 == WS {
		event.Date = lit1
		err = p.parseSingle(event)
		events = append(events, event)
	}

	if len(events) == 0 || (err != nil && err.Error() != "EOF") {
		return []*Event{nil}, err
	}

	return events, nil
}

func (p *Parser) ParseAll() ([]*Event, error) {
	var events []*Event

	for {
		event, err := p.Parse()

		if err != nil {
			return events, err
		} else if slices.Contains(event, nil) {
			break
		}

		events = append(events, event...)
	}

	return events, nil
}
