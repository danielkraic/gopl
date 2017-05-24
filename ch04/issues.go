package main

import (
	"github.com/danielkraic/gopl/ch04/github"
	"os"
	"fmt"
)

func main() {
	res, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(10)
	}

	fmt.Printf("count: %d\n", res.TotalCount)
	for _, issue := range res.Items {
		fmt.Printf("%-5d %9.9s %.55s\n", issue.Number, issue.User.Login, issue.Title)
	}
}
