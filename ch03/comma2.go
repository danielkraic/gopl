package main

import (
	"fmt"
	"bytes"
	"strings"
)

func comma(s string) string {
	if len(s) == 0 {
		return s
	}

	var prefix, postfix string

	// get prefix
	if s[0] == '+' || s[0] == '-' {
		prefix = string(s[0])
		s = s[1:]
	}

	// get postfix
	dot := strings.LastIndex(s, ".")
	if dot != -1 {
		postfix = s[dot:]
		s = s[:dot]
	} else {
		e := strings.LastIndex(s, "e")
		if e != -1 {
			postfix = s[e:]
			s = s[:e]
		} else {
			e := strings.LastIndex(s, "E")
			if e != -1 {
				postfix = s[e:]
				s = s[:e]
			}
		}
	}

	rem := len(s) % 3
	start := s[:rem]
	s = s[rem:]

	var buf bytes.Buffer
	buf.WriteString(prefix)
	buf.WriteString(start)
	if len(start) > 0 && len(s) > 0 {
		buf.WriteRune(',')
	}

	for i, ch := range s {
		if i != 0 && i % 3 == 0 {
			buf.WriteRune(',')
		}
		buf.WriteRune(ch)
	}
	buf.WriteString(postfix)

	return buf.String()
}

func printComma(s string) {
	fmt.Printf("%-10s : %-10s\n", s, comma(s))
}

func main() {
	printComma("")
	printComma("1")
	printComma("12")
	printComma("123")
	printComma("1234")
	printComma("12345")
	printComma("123456")
	printComma("1234567")

	printComma("1.2")
	printComma("1.23")
	printComma("1.234")
	printComma("1.2345")
	printComma("12.1234")
	printComma("123.1234")
	printComma("1234.1234")
	printComma("1234567.1234")

	printComma("1.2e4")
	printComma("1.23e4")
	printComma("1.234e4")
	printComma("1.2345e4")
	printComma("12.1234e4")
	printComma("123.1234e4")
	printComma("1234.1234e4")
	printComma("1234567.1234e4")

	printComma("1.2E4")
	printComma("1.23E4")
	printComma("1.234E4")
	printComma("1.2345E4")
	printComma("12.1234E4")
	printComma("123.1234E4")
	printComma("1234.1234E4")
	printComma("1234567.1234E4")
}
