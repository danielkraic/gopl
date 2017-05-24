package main

import "fmt"

var graph = make(map[string]map[string]bool)

func addEdge(from, to string) {
	edges := graph[from]
	if edges == nil {
		edges = make(map[string]bool)
		graph[from] = edges
	}

	edges[to] = true
}

func hasEdge(from, to string) bool {
	return graph[from][to]
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
