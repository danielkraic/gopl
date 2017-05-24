package main

import (
	"fmt"
	"sort"
)

func main() {
	m1 := make(map[string]int)

	m2 := map[string]int{
		"a": 1,
		"b": 2,
	}

	m2["c"] = 5
	m2["d"] += 10
	m2["e"]++

	delete(m2, "a")

	fmt.Printf("%v\n", m1)
	fmt.Printf("%v\n", m2)

	//m2["somekey"] = 10
	if val, ok := m2["somekey"]; ok {
		fmt.Printf("somekey val: %v\n", val)
	} else {
		fmt.Printf("no value for somekey\n")
	}

	keys := make([]string, 0, len(m2))
	for k := range m2 {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	for _, k := range keys {
		fmt.Printf("%v: %v\n", k, m2[k])
	}
}
