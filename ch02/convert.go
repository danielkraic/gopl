package main

import (
	"os"
	"bufio"
	"fmt"
	"strconv"
	"github.com/danielkraic/gopl/ch02/tempconv"
)

func main() {
	if len(os.Args) > 1 {
		for _, arg := range os.Args[1:] {
			processValue(arg)
		}
	} else {
		reader := bufio.NewScanner(os.Stdin)
		for reader.Scan() {
			processValue(reader.Text())
		}
		if err := reader.Err(); err != nil {
			fmt.Fprint(os.Stderr, "Failed to read from STDIN. Error: %s\n", err)
			os.Exit(1)
		}
	}
}

func processValue(value string) {
	number, err := strconv.ParseFloat(value, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr,"Failed to parse float from string value: '%s'\n", err)
	} else {
		c := tempconv.Celsius(number)
		fmt.Printf("C: %s == %s\n", c, tempconv.CToF(c))

		f := tempconv.Farenheit(number)
		fmt.Printf("F: %s == %s\n", f, tempconv.FToC(f))
	}
}