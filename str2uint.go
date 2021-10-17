package main

import (
	"log"
	"strconv"
)

func main() {
	s := "8080"
	u, err := strconv.ParseUint(s, 10, 16)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(u)
}