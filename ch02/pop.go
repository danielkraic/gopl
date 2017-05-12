package main

import (
	"fmt"
)

func main() {
	//var pc[16] byte
	//for i := range pc {
	//	pc[i] = pc[i/2] + byte(i&1)
	//
	//	fmt.Printf("%b (%d): %v + %b == %v\n", i, i, pc[i/2], byte(i&1), pc[i/2] + byte(i&1))
	//}

	x := 1000000
	fmt.Printf("%b\n\n", x)
	fmt.Printf("%b\n", byte(x>>(0*8)))
	fmt.Printf("%b\n", byte(x>>(1*8)))
	fmt.Printf("%b\n", byte(x>>(2*8)))
	fmt.Printf("%b\n", byte(x>>(3*8)))
	fmt.Printf("%b\n", byte(x>>(4*8)))
	fmt.Printf("%b\n", byte(x>>(5*8)))
	fmt.Printf("%b\n", byte(x>>(6*8)))
	fmt.Printf("%b\n", byte(x>>(7*8)))
}