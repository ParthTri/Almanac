package parser

import (
	"errors"
	"slices"
	"strings"
	"testing"
)

func prettyToken(tok Token) (string, error) {
	tokens := map[Token]string{
		ILLIGAL:     "ILLIGAL",
		EOF:         "EOF",
		EOL:         "EOL",
		WS:          "WS",
		TAG:         "TAG",
		DESCRIPTION: "DESCRIPTION",
		WORD:        "WORD",
		DATE:        "DATE",
		TIME:        "TIME",
		DASH:        "DASH",
	}

	out := tokens[tok]

	if out == "" {
		return "", errors.New("Unkown Token")
	}

	return out, nil
}

func TestScanLetter(t *testing.T) {
	input := "This is my input with + and *"
	reader := strings.NewReader(input)
	scanner := NewScanner(reader)

	tok, val := scanner.scanLetter()

	t.Log(tok)
	t.Log(val)

	if val != "This" {
		t.Errorf("Got wrong value")
	}
}

func TestScanWhiteSpace(t *testing.T) {
	input := "    "
	reader := strings.NewReader(input)
	scanner := NewScanner(reader)

	tok, val := scanner.scanWhitespace()

	t.Log(tok)
	t.Log(val)

	if tok != 3 {
		t.Errorf("Wrong token found %v. Expected WS 3", tok)
	}
}

func TestScan(t *testing.T) {
	input := "This is my input +hello and *"
	reader := strings.NewReader(input)
	scanner := NewScanner(reader)

	var tokens []Token

	for {
		tok, _ := scanner.Scan()

		if tok == EOF {
			break
		}

		t.Log(tok)
		out, err := prettyToken(tok)
		if err != nil {
			t.Errorf("Unkown token found %v", tok)
			return
		}
		t.Logf("%v ", out)

		tokens = append(tokens, tok)
	}

	t.Log(tokens)
	if !slices.Equal(tokens, []Token{WORD, WS, WORD, WS, WORD, WS, WORD, WS, TAG, WORD, WS, WORD, WS, DESCRIPTION}) {
		t.Errorf("Tokens do not match, got %v", tokens)
	}
}

func TestNewLine(t *testing.T) {
	input := "* Hello\n+nice"
	reader := strings.NewReader(input)
	scanner := NewScanner(reader)

	var tokens []Token

	for {
		tok, _ := scanner.Scan()

		if tok == EOF {
			break
		}

		out, err := prettyToken(tok)
		if err != nil {
			t.Errorf("Unkown token found %v", tok)
			return
		}
		t.Logf("%v ", out)
		tokens = append(tokens, tok)
	}

	if !slices.Equal(tokens, []Token{DESCRIPTION, WS, WORD, EOL, TAG, WORD}) {
		t.Errorf("Tokens do not match, got %v", tokens)
	}
}

func TestScanYear(t *testing.T) {
	input := " 2025-12-12\n"
	reader := strings.NewReader(input)
	scanner := NewScanner(reader)

	var tokens []Token

	for {
		tok, _ := scanner.Scan()

		if tok == EOF {
			break
		}

		out, err := prettyToken(tok)
		if err != nil {
			t.Errorf("Unkown token found %v", tok)
			return
		}
		t.Logf("%v ", out)
		tokens = append(tokens, tok)
	}

	if !slices.Equal(tokens, []Token{WS, DATE, EOL}) {
		t.Errorf("Tokens do not match, got %v", tokens)
	}
}

func TestScanTime(t *testing.T) {
	input := "time 14:00\n\n"
	reader := strings.NewReader(input)
	scanner := NewScanner(reader)

	var tokens []Token

	for {
		tok, lit := scanner.Scan()

		if tok == EOF {
			break
		} else if tok == ILLIGAL {
			t.Errorf("Found illegal token '%v'", lit)
			return
		}

		out, err := prettyToken(tok)
		if err != nil {
			t.Errorf("Unkown token found %v", out)
			return
		}

		t.Log(out)

		tokens = append(tokens, tok)
	}

	if !slices.Equal(tokens, []Token{WORD, WS, TIME, EOL, EOL}) {
		t.Errorf("Tokens do not match, got %v", tokens)
	}
}

func TestScanTimeRange(t *testing.T) {
	input := "time  14:00-16:00\n\n"
	reader := strings.NewReader(input)
	scanner := NewScanner(reader)

	var tokens []Token

	for {
		tok, lit := scanner.Scan()

		if tok == EOF {
			break
		} else if tok == ILLIGAL {
			t.Errorf("Found illegal token '%v'", lit)
			return
		}

		out, err := prettyToken(tok)
		if err != nil {
			t.Errorf("Unkown token found %v", out)
			return
		}

		t.Log(out)

		tokens = append(tokens, tok)
	}

	if !slices.Equal(tokens, []Token{WORD, WS, TIME, DASH, TIME, EOL, EOL}) {
		t.Errorf("Tokens do not match, got %v", tokens)
	}
}
