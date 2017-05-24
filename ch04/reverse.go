package main

import "fmt"

func reverse(s []int) {
	for i, j := 0, len(s) - 1; i < j; i, j = i + 1, j - 1 {
		s[i], s[j] = s[j], s[i]
	}
}

func printReverse(s []int) {
	fmt.Printf("%v\n", s)
	reverse(s)
	fmt.Printf("%v\n", s)
}

func main() {
	a1 := [...]int{}
	printReverse(a1[:])

	a2 := [...]int{1}
	printReverse(a2[:])

	a3 := [...]int{1,2}
	printReverse(a3[:])

	a4 := [...]int{1,2,3}
	printReverse(a4[:])

	a5 := [...]int{1,2,3, 4}
	printReverse(a5[:])

	a := make([]string, 0, 2)
	a = append(a, "a")
	a = append(a, "b", "e")
	a = append(a, "f", "g")
	fmt.Printf("%v, len %d cap %d\n", a, len(a), cap(a))

}
