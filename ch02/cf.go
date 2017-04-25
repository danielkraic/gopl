package main

import (
	"log"
	"os"
	"strconv"

	"fmt"

	"github.com/danielkraic/gopl/ch02/tempconv"
)

func main() {
	for _, val := range os.Args[1:] {
		num, err := strconv.ParseFloat(val, 64)
		if err != nil {
			log.Fatalf("Invalid float value '%s'\n", val)
		}

		c := tempconv.Celsius(num)
		f := tempconv.Farenheit(num)

		fmt.Printf("%v == %s\n", f, tempconv.FToC(f))
		fmt.Printf("%v == %s\n", c, tempconv.CToF(c))
	}
}
