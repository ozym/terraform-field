package ros

import (
	"bufio"
	"strconv"
	"strings"
	"text/scanner"
	"unicode"
)

// Scan the Flags line providing a list of key value pairs.
// e.g "Flags: X - disabled, I - invalid, D - dynamic"
func scanFlags(line string) map[string]string {
	var s scanner.Scanner
	s.Init(strings.NewReader(line))

	s.Whitespace = 1<<'\t' | 1<<'-' | 1<<'\r' | 1<<' ' | 1<<','
	s.IsIdentRune = func(ch rune, i int) bool {
		return ch == '*' || unicode.IsLetter(ch) || unicode.IsDigit(ch)
	}
	s.Mode = scanner.ScanIdents

	flags := make(map[string]string)

	var tok rune
	for {
		var key string
		if tok = s.Scan(); tok == scanner.EOF || tok == '\n' {
			break
		}
		key = s.TokenText()
		if tok = s.Scan(); tok == scanner.EOF || tok == '\n' {
			break
		}
		flags[key] = s.TokenText()
	}

	return flags
}

// Scan any actual comment lines that begin with ';;;'
func scanComment(s *scanner.Scanner) string {
	ws := s.Whitespace
	defer func() {
		s.Whitespace = ws
	}()
	s.Whitespace = 1 << '\r'

	comment := ""

	var tok rune
	for tok != scanner.EOF {
		tok = s.Scan()
		if tok == '\n' {
			break
		}
		comment = comment + s.TokenText()
	}
	return comment
}

func scanLine(s *scanner.Scanner) string {
	ws := s.Whitespace
	defer func() {
		s.Whitespace = ws
	}()
	s.Whitespace = 1 << '\r'

	line := s.TokenText()

	var tok rune
	for tok != scanner.EOF {
		if s.Peek() == '\n' {
			break
		}
		if tok = s.Scan(); tok == scanner.EOF {
			continue
		}
		line = line + s.TokenText()
	}
	return line
}

// assumes the current position has been processed
func scanOpts(s *scanner.Scanner) (string, string, string) {

	ws := s.Whitespace
	mo := s.Mode
	id := s.IsIdentRune

	defer func() {
		s.Whitespace = ws
		s.Mode = mo
		s.IsIdentRune = id
	}()

	// only expect letters (at least for now)
	s.Whitespace = 1<<'\t' | 1<<'\r' | 1<<' ' | 1<<','
	s.IsIdentRune = func(ch rune, i int) bool {
		return ch == '-' || ch == ';' || ch == '*' || unicode.IsLetter(ch) || unicode.IsDigit(ch)
	}
	s.Mode = scanner.ScanIdents

	var opts, comment, line string

	var tok rune
	for tok != scanner.EOF && tok != '\n' {
		if tok = s.Scan(); tok == scanner.EOF || tok == '\n' {
			continue
		}
		if tok == scanner.Ident {
			if s.TokenText() == ";;;" {
				// skip the current commas
				if tok = s.Scan(); tok == scanner.EOF {
					continue
				}
				// recover the comment text
				comment = scanLine(s)
			} else if s.Peek() != '=' {
				opts += s.TokenText()
			} else {
				line += s.TokenText()
			}
		} else {
			line += scanLine(s)
		}
	}

	return opts, comment, line
}

// scan a script, i.e. "{/ip address set [find address="10.54.242.1/28" ] disabled=no}"
func scanRaw(s *scanner.Scanner) string {
	var res string

	mode := s.Mode
	ws := s.Whitespace
	defer func() {
		s.Mode = mode
		s.Whitespace = ws
	}()

	s.Mode = 0
	s.Whitespace = 0

	// have one already
	bracket := 1

	var tok rune
	for tok != scanner.EOF {
		if tok = s.Scan(); tok == scanner.EOF {
			continue
		}
		if tok == '{' {
			bracket++
		} else if tok == '}' {
			bracket--
		}
		res += string(tok)

		if bracket == 0 {
			break
		}
	}

	return res
}

func isIdent(ch rune) bool {
	for _, i := range ":.;/-,_[]" {
		if ch == i {
			return true
		}
	}
	return false
}

func scanKeyValues(line string) map[string]string {
	var key string
	res := make(map[string]string)

	var s scanner.Scanner
	s.Init(strings.NewReader(line))

	s.Mode = scanner.ScanIdents | scanner.ScanStrings
	s.Whitespace = 1<<'\t' | 1<<'\r' | 1<<'\n' | 1<<' '
	s.IsIdentRune = func(ch rune, i int) bool {
		return isIdent(ch) || unicode.IsLetter(ch) || unicode.IsDigit(ch)
	}

	var tok rune
	for tok != scanner.EOF {
		if tok = s.Scan(); tok == scanner.EOF {
			continue
		}

		if s.Peek() == '=' {
			key = s.TokenText()
		} else if s.TokenText() == "{" {
			res[key] += s.TokenText() + scanRaw(&s)
		} else if s.TokenText() != "=" {
			u, err := strconv.Unquote(s.TokenText())
			if err != nil {
				u = s.TokenText()
			}
			if _, ok := res[key]; ok {
				res[key] = res[key] + " " + u
			} else {
				res[key] = u
			}
		}
	}

	return res
}

// Scan a set of numbered items, this may be preceded with a Flags line.
func ScanNumberedItemList(results string) ([]map[string]string, error) {
	var flags map[string]string

	reader := bufio.NewReader(strings.NewReader(results))

	var lines []string
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		lines = append(lines, string(line))
	}

	// precondition the input lines ...
	for i := 1; i < len(lines); i++ {
		switch {
		case strings.HasPrefix(strings.TrimLeftFunc(lines[i-1], unicode.IsSpace), "#"):
			// remove any general comments
			lines = append(lines[:i-1], lines[i:]...)
		case strings.HasPrefix(strings.TrimLeftFunc(lines[i-1], unicode.IsSpace), "Flags:"):
			// recover the flags line
			flags = scanFlags(strings.Replace(lines[i-1], "Flags:", "", -1))
			// cut out the previous line
			lines = append(lines[:i-1], lines[i:]...)
		case strings.HasSuffix(strings.TrimRightFunc(lines[i-1], unicode.IsSpace), ","):
			// update the previous line
			lines[i-1] = strings.TrimRightFunc(lines[i-1], unicode.IsSpace) + strings.TrimSpace(lines[i])
			// cut out the current line
			if i < len(lines)-1 {
				lines = append(lines[:i], lines[i+1:]...)
			} else {
				lines = lines[:i]
			}
		}
	}

	var s scanner.Scanner
	s.Init(strings.NewReader(strings.Join(lines, "\n")))

	s.Mode = scanner.ScanInts | scanner.ScanIdents
	s.Whitespace = 1<<'\t' | 1<<'\r' | 1<<'\n' | 1<<' '
	s.IsIdentRune = func(ch rune, i int) bool {
		return ch == ':' || ch == ';' || ch == '-' || unicode.IsLetter(ch)
	}

	var list []map[string]string

	var number, opts, comment, line string

	var tok rune
	for tok != scanner.EOF {
		if tok = s.Scan(); tok == scanner.EOF {
			continue
		}
		if tok == scanner.Int {
			if number != "" && line != "" {
				ans := map[string]string{
					"number":  number,
					"comment": comment,
				}
				lookup := make(map[string]bool)
				for _, o := range opts {
					lookup[string(o)] = true
				}
				for k, v := range flags {
					if _, ok := lookup[k]; ok {
						ans[v] = "yes"
					} else {
						ans[v] = "no"
					}
				}
				for k, v := range scanKeyValues(line) {
					ans[k] = v
				}
				list = append(list, ans)
			}
			number = s.TokenText()
			opts, comment, line = scanOpts(&s)
		} else {
			line += " " + scanLine(&s)
		}
	}

	if line != "" {
		if number != "" && line != "" {
			ans := map[string]string{
				"number":  number,
				"comment": comment,
			}
			lookup := make(map[string]bool)
			for _, o := range opts {
				lookup[string(o)] = true
			}
			for k, v := range flags {
				if _, ok := lookup[k]; ok {
					ans[v] = "yes"
				} else {
					ans[v] = "no"
				}
			}
			for k, v := range scanKeyValues(line) {
				ans[k] = v
			}
			list = append(list, ans)
		}
	}

	return list, nil
}

// Take the first entry, usually after applying a filter.
func ScanFirstNumberedItemList(results string) (map[string]string, error) {
	list, err := ScanNumberedItemList(results)
	if err != nil {
		return nil, err
	}
	if len(list) > 0 {
		return list[0], nil
	}
	return nil, nil
}
