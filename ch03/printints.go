package main

import (
	"bytes"
	"fmt"
)

func intsToString(ints []int) string {
	var buf bytes.Buffer

	_, err := buf.WriteRune('[')
	if err != nil {
		panic(err)
	}

	for i, val := range ints {
		if i > 0 {
			_, err = buf.WriteString(", ")
			if err != nil {
				panic(err)
			}
		}
		fmt.Fprintf(&buf, "%d", val)
	}

	_, err = buf.WriteRune(']')
	if err != nil {
		panic(err)
	}

	return buf.String()
}

func printInts(ints []int) {
	fmt.Printf("%v : %s\n", ints, intsToString(ints))
}

func main() {
	printInts([]int{})
	printInts([]int{1})
	printInts([]int{1,2})
	printInts([]int{1,2,3})
	printInts([]int{1,2,3,4})
	printInts([]int{1,2,3,4,5})
}
