package parser

import (
	"bufio"
	"bytes"
	"io"
	"unicode"
)

type Token int

const (
	ILLIGAL Token = iota
	EOF
	EOL // End of Line
	WS  // White space

	// MAIN FEATURES
	TAG         // +
	DESCRIPTION // *
	DASH        // - For time range

	WORD // Characters
	DATE // Date YYYY-MM-DD
	TIME // Time HH:MM

)

type Scanner struct {
	r   *bufio.Reader
	buf []rune
}

func isWhitespace(ch rune) bool {
	return unicode.IsSpace(ch)
}

func isLetter(ch rune) bool {
	return unicode.IsLetter(ch)
}

func isDigit(ch rune) bool {
	return unicode.IsDigit(ch)
}

func NewScanner(r io.Reader) *Scanner {
	return &Scanner{r: bufio.NewReader(r)}
}

func (s *Scanner) read() rune {
	// If we have runes in the unread buffer, pop from there.
	if len(s.buf) > 0 {
		ch := s.buf[len(s.buf)-1]
		s.buf = s.buf[:len(s.buf)-1]
		return ch
	}

	ch, _, err := s.r.ReadRune()
	if err != nil {
		return eof
	}

	return ch
}

func (s *Scanner) unread(ch rune) {
	s.buf = append(s.buf, ch)
}

func (s *Scanner) Scan() (tok Token, lit string) {
	ch := s.read()

	switch ch {
	case eof:
		return EOF, ""
	case '+':
		return TAG, string(ch)
	case '*':
		return DESCRIPTION, string(ch)
	case '\n':
		return EOL, string(ch)
	case '-':
		return DASH, string(ch)
	}

	if isWhitespace(ch) {
		s.unread(ch)
		return s.scanWhitespace()
	} else if isDigit(ch) {
		s.unread(ch)
		tok, lit := s.scanDate()

		if tok == ILLIGAL {
			for i := len(lit) - 1; i >= 0; i-- {
				s.unread(rune(lit[i]))
			}
			tok, lit = s.scanTime()
		}

		return tok, lit
	} else if isLetter(ch) {
		s.unread(ch)
		return s.scanLetter()
	}

	return ILLIGAL, string(ch)
}

func (s *Scanner) scanWhitespace() (tok Token, lit string) {
	// Create a buffer an read the current character into it
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isWhitespace(ch) {
			s.unread(ch)
			break
		} else {
			buf.WriteRune(ch)
		}
	}
	return WS, buf.String()
}

func (s *Scanner) scanLetter() (tok Token, lit string) {
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isLetter(ch) && !isDigit(ch) && ch != '_' {
			s.unread(ch)
			break
		} else {
			_, _ = buf.WriteRune(ch)
		}
	}

	return WORD, buf.String()
}

func (s *Scanner) scanDate() (tok Token, lit string) {
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isDigit(ch) && ch != '-' {
			s.unread(ch)
			break
		} else {
			_, _ = buf.WriteRune(ch)
		}
	}

	if buf.Len() != 20 {
		return ILLIGAL, buf.String()
	}

	return DATE, buf.String()
}

func (s *Scanner) scanTime() (tok Token, lit string) {
	var buf bytes.Buffer = *bytes.NewBuffer(make([]byte, 5))
	buf.WriteRune(s.read())

	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isDigit(ch) && ch != ':' {
			s.unread()
			break
		} else {
			_, _ = buf.WriteRune(ch)
		}
	}

	return TIME, buf.String()
}
