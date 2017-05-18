package main

import "fmt"

func zero(ptr *[32]byte) {
	for i := range ptr {
		ptr[i] = 0
	}
}

func main() {
	a1 := [32]byte{2: byte(1), 31: byte(1)}

	fmt.Printf("%T %v\n", a1, a1)
	zero(&a1)
	fmt.Printf("%T %v\n", a1, a1)
}
