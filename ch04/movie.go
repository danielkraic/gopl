package main

import (
	"fmt"
	"encoding/json"
	"os"
)

type Movie struct {
	Title string
	Year int `json:"released"`
	Color bool `json:"color,omitempty"`
	Actors []string
}

func main() {
	movies := []Movie{
		{Title: "Cassablanca", Year: 1942, Color: false,
		Actors: []string{"HB", "IB"}},
		{Title: "Cool Hand Luke", Year: 1967, Color: true,
			Actors: []string{"PN"}},
	}

	fmt.Printf("%v\n", movies)

	data1, err1 := json.Marshal(&movies)
	if err1 != nil {
		fmt.Fprintf(os.Stderr, "Json marshall failed: %s\n", err1)
		os.Exit(10)
	}
	fmt.Printf("%s\n", data1)

	data2, err2 := json.MarshalIndent(&movies, "", " ")
	if err2 != nil {
		fmt.Fprintf(os.Stderr, "Json marshall failed: %s\n", err2)
		os.Exit(11)
	}
	fmt.Printf("%s\n", data2)

	var data3 []Movie
	err3 := json.Unmarshal(data2, &data3)
	if err3 != nil {
		fmt.Fprintf(os.Stderr, "Error while unmarshalling: %\n", err3)
		os.Exit(12)
	}
	fmt.Printf("%v\n", data3)
}
