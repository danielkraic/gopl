package main

import (
	"fmt"
	"unicode/utf8"
	"bufio"
	"os"
	"io"
	"unicode"
)

func main() {
	counts := make(map[rune]int)
	var utflen [utf8.UTFMax + 1]int
	invalid := 0
	category := make(map[string]int)

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune()

		if err == io.EOF {
			break;
		}

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			break
		}

		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}

		counts[r]++
		utflen[n]++

		if unicode.IsControl(r) {
			category["CONTROL"]++
		}
		if unicode.IsDigit(r) {
			category["DIGIT"]++
		}
		if unicode.IsLetter(r) {
			category["LETTER"]++
		}
		if unicode.IsSpace(r) {
			category["SPACE"]++
		}
		if unicode.IsPunct(r) {
			category["PUNCT"]++
		}
	}

	fmt.Printf("rune\tcount\t\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}

	fmt.Printf("nlen\tcount\n")
	for i, n := range utflen {
		fmt.Printf("%d\t%d\n", i, n)
	}

	fmt.Printf("category\tcount\n")
	for c, n := range category {
		fmt.Printf("%s\t%d\n", c, n)
	}

	if invalid > 0 {
		fmt.Printf("Invalid count: %d\n", invalid)
	}
}
