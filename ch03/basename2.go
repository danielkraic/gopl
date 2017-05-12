package main

import (
	"fmt"
	"strings"
)

func basename(path string) string {
	//for i := len(path) - 1; i >= 0; i-- {
	//	if path[i] == '/' {
	//		path = path[i+1:]
	//		break
	//	}
	//}

	slash := strings.LastIndex(path, "/")
	if slash != -1 {
		path = path[slash+1:]
	}

	dot := strings.LastIndex(path, ".")
	if dot != -1 {
		path = path[:dot]
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
