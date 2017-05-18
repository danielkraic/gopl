package main

import "fmt"

type Currency int

const (
	EUR Currency = iota
	USD
	GBP
)

func main() {
	names := [...]string{EUR: "euro", USD: "usd", GBP: "gbp"}
	fmt.Printf("%s\n", names[EUR])
	fmt.Printf("%s\n", names[USD])
	fmt.Printf("%s\n", names[GBP])
	fmt.Printf("%v\n", names)
}
