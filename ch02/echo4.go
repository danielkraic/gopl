package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	var n = flag.Bool("n", false, "omit trailing newline")
	var sep = flag.String("s", " ", "separator")

	flag.Parse()

	fmt.Print(strings.Join(os.Args[1:], *sep))
	if !*n {
		fmt.Println("")
	}
}
