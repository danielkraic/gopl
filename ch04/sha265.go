package main

import (
	"bufio"
	"os"
	"crypto/sha256"
	"fmt"
)

func main() {
	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		sum := sha256.Sum256([]byte(in.Text()))
		fmt.Printf("%x\n", sum)
	}
}
