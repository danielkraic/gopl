package main

import "fmt"

func main() {
	var x uint8 = 1<<5 | 1 << 1
	var y uint8 = 1<<2| 1 << 1

	fmt.Printf("x: %08b set{1,5}\n", x)
	fmt.Printf("y: %08b set{1,2}\n", y)

	fmt.Printf("x&y: %08b intersection\n", x&y)
	fmt.Printf("x|y: %08b union\n", x|y)
	fmt.Printf("x^y: %08b symmetric difference\n", x^y)
	fmt.Printf("x&^y: %08b difference\n", x&^y)

	for i := uint(0); i < 8; i++ {
		if x&(1<<i) != 0 {
			fmt.Printf("%d in %d\n", i, x)
		}
	}
}

