package main

import "fmt"

func makeRuneMap(s string) map[rune]int {
	runes := make(map[rune]int)
	for _, r := range s {
		runes[r]++
	}
	return runes
}

func isAnagram(s1, s2 string) bool {
	if len(s1) == 0 || len(s1) != len(s2) {
		return false
	}

	runes1 := makeRuneMap(s1)
	runes2 := makeRuneMap(s2)

	for r, val1 := range runes1 {
		val2, found := runes2[r]
		if !found || val1 != val2 {
			return false
		}
	}

	for r, val2 := range runes2 {
		val1, found := runes1[r]
		if !found || val1 != val2 {
			return false
		}
	}

	return true
}

func printIsAnagram(s1, s2 string) {
	fmt.Printf("%v :: '%s' ::' %s'\n", isAnagram(s1, s2), s1, s2)
}

func main() {
	printIsAnagram("", "")
	printIsAnagram("", "a")
	printIsAnagram("a", "")
	printIsAnagram("a", "a")
	printIsAnagram("aa", "a")
	printIsAnagram("a", "aa")
	printIsAnagram("ab", "ab")
	printIsAnagram("ab", "ba")
	printIsAnagram("ab", "baa")
	printIsAnagram("abb", "ba")
	printIsAnagram("ab", "ac")
	printIsAnagram("ab", "ac")
}
