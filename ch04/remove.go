package main

import "fmt"

func remove(strings []string, i int) []string {
	copy(strings[i:], strings[i+1:])
	return strings[:len(strings)-1]
}

func main() {
	a := []string{"a", "b", "c", "d", "e"}
	fmt.Printf("%v\n", a)
	a = remove(a, 3)
	fmt.Printf("%v\n", a)
}
