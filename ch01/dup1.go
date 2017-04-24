package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	data := make(map[string]int)
	for input.Scan() {
		data[input.Text()]++
	}

	for line, count := range data {
		fmt.Printf("%-5d: %s\n", count, line)
	}
}
