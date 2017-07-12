package ros

import (
	"strings"
	"text/scanner"
	"unicode"
)

type Token struct {
	Type rune
	Peek rune
	Text string
}

func (t Token) String() string {
	return t.Text
}

type Scanner struct {
	scanner.Scanner
	tok Token
}

func NewScanner(str string) Scanner {
	s := Scanner{Scanner: scanner.Scanner{}}
	s.Init(strings.NewReader(str))
	return s
}

func (s *Scanner) Token() Token {
	return s.tok
}
func (s *Scanner) String() string {
	return s.tok.Text
}
func (s *Scanner) Type() rune {
	return s.tok.Type
}

// Next iterates over all tokens. Retrieve the most recent token with Token(). It returns false once it reaches token.EOF.
func (s *Scanner) Next() bool {
	s.tok.Type = s.Scan()
	if s.tok.Type == scanner.EOF {
		return false
	}
	s.tok.Text = s.TokenText()
	s.tok.Peek = s.Peek()
	return true
}

// <ws><key>:[<value>...]<nl>
func ScanItems(results string) (map[string]string, error) {
	items := make(map[string]string)

	s := NewScanner(results)
	s.Mode = scanner.ScanIdents
	s.Whitespace = 1<<'\t' | 1<<'\r' | 1<<'"'
	s.IsIdentRune = func(ch rune, i int) bool {
		return ch == '.' || ch == '-' || unicode.IsLetter(ch) || unicode.IsDigit(ch)
	}
	for s.Next() {
		var key, value string
		for s.Next() {
			if s.Type() == ':' {
				break
			}
			key = s.String()
		}
		for s.Next() {
			if s.Type() == '\n' {
				break
			}
			value = value + s.String()
		}
		if len(strings.TrimSpace(key)) > 0 {
			items[strings.TrimSpace(key)] = strings.TrimSpace(value)
		}
	}

	return items, nil
}
