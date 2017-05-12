package main

import (
	"fmt"
)

func basename(path string) string {
	for i := len(path) - 1; i >= 0; i-- {
		if path[i] == '/' {
			path = path[i+1:]
			break
		}
	}

	for i := len(path) - 1; i >= 0; i-- {
		if path[i] == '.' {
			path = path[:i]
			break
		}
	}

	return path
}

func printBasename(path string) {
	fmt.Printf("%-10s:\t\t%s\n", path, basename(path))
}

func main() {
	printBasename("")
	printBasename("/")
	printBasename("a")
	printBasename("a/")
	printBasename("/a/")
	printBasename("/a/")
	printBasename("/a")
	printBasename("/a.b")
	printBasename("/a.b")
}
