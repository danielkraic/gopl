package main

import "fmt"

func nonempty(strings []string) []string {
	i := 0
	for _, val := range strings {
		if val != "" {
			strings[i] = val
			i++
		}
	}
	return strings[:i]
}

func print(strings []string) {
	len1 := len(strings)
	strings = nonempty(strings)
	len2 := len(strings)

	fmt.Printf("%d -> %d: %v\n", len1, len2, strings)
}

func main() {
	print([]string{})
	print([]string{""})
	print([]string{"",""})
	print([]string{"a",""})
	print([]string{"","a"})
	print([]string{"","a", ""})
	print([]string{"","a", "a"})
	print([]string{"a","a", ""})
	print([]string{"a","", "a"})
	print([]string{"a","a", "a"})
}
