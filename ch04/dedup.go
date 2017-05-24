package main

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	m := make(map[string]bool)
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		str := s.Text()
		if _, ok := m[str]; !ok {
			m[str] = true
			fmt.Println(str)
		}
	}
	if err := s.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Err: %v\n", err)
	}
}
