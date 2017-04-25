package main

import "fmt"

func main() {
	freezingF, boilingF := 32.0, 212.0

	fmt.Printf("Freezing: %v°C == %v°F\n", fToC(freezingF), freezingF)
	fmt.Printf("Freezing: %v°C == %v°F\n", fToC(boilingF), boilingF)
}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}
