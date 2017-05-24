package main

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	m := make(map[string]bool)
	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)
	for in.Scan() {
		val := in.Text()
		if _, ok := m[val]; !ok {
			m[val] = true
			fmt.Println(val)
		}
	}
}
