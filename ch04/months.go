package main

import "fmt"

const (
	JAN = 1 + iota
	FEB
	MAR
	APR
	MAY
	JUN
	JUL
	AUG
	SEP
	OCT
	NOV
	DEC
)

func main() {
	months := [...]string{
		JAN: "jan",
		FEB: "feb",
		MAR: "mar",
		APR: "apr",
		MAY: "may",
		JUN: "jun",
		JUL: "jul",
		AUG: "aug",
		SEP: "sep",
		OCT: "oct",
		NOV: "noc",
		DEC: "dec",
	}

	Q1 := months[1:4]
	Q2 := months[4:7]
	Q3 := months[7:10]
	Q4 := months[10:]

	fmt.Printf("%T %v\n", months, months)
	for i, val := range months {
		fmt.Printf("%d %v\n", i, val)
	}
	fmt.Printf("%T %v\n", Q1, Q1)
	fmt.Printf("%T %v\n", Q2, Q2)
	fmt.Printf("%T %v\n", Q3, Q3)
	fmt.Printf("%T %v\n", Q4, Q4)
}
