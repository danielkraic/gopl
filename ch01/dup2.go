package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	counts := make(map[string]int)
	if len(os.Args[1:]) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, filePath := range os.Args[1:] {
			file, err := os.Open(filePath)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Unable to open file '%s'\n", filePath)
				continue
			}
			countLines(file, counts)
		}
	}

	for line, count := range counts {
		fmt.Printf("%-5d: %s\n", count, line)
	}
}

func countLines(reader io.Reader, counts map[string]int) {
	input := bufio.NewScanner(reader)
	for input.Scan() {
		counts[input.Text()]++
	}
}
