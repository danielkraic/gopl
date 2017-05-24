package main

import "fmt"

var graph = make(map[string]map[string]bool)

func addEdge(from, to string) {
	if _, ok := graph[from]; !ok {
		graph[from] = make(map[string]bool)
	}

	graph[from][to] = true
}

func hasEdge(from, to string) bool {
	if _, ok := graph[from]; ok {
		if val, ok := graph[from][to]; ok {
			return val
		}
	}
	return false
}

func main() {
	fmt.Printf("%q\n", graph)

	fmt.Printf("has: %v\n", hasEdge("a", "b"))
	addEdge("a", "b")
	fmt.Printf("has: %v\n", hasEdge("a", "b"))
	fmt.Printf("%q\n", graph)


	addEdge("a", "c")
	fmt.Printf("has: %v\n", hasEdge("a", "c"))
	fmt.Printf("%q\n", graph)
}
