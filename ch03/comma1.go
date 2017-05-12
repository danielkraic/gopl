package main

import "fmt"

func comma(s string) string {
	if len(s) <= 3 {
		return s
	}

	return comma(s[:len(s)-3]) + "," + s[len(s)-3:]
}

func printComma(s string) {
	fmt.Printf("%-10s : %-10s\n", s, comma(s))
}

func main() {
	printComma("")
	printComma("a")
	printComma("ab")
	printComma("abc")
	printComma("abcd")
	printComma("abcde")
	printComma("abcdef")
	printComma("abcdefg")
}
