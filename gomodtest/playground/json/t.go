package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title string
	Year int
	Color bool
	Actors []string
}

var movies = []Movie{
	{
		Title:  "aa",
		Year:   1988,
		Color:  false,
		Actors: []string{"aabb", "abb"},
	},
	{
		Title:  "bb",
		Year:   1966,
		Color:  true,
		Actors: []string{"123", "456"},
	},
	{
		Title:  "cc",
		Year:   1999,
		Color:  false,
		Actors: []string{"da66", "87gr"},
	},
}

func main() {
	data, err := json.Marshal(movies)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))

	data1, err := json.MarshalIndent(movies, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data1))
}
