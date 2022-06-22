package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := "70.5"
	i, err := strconv.Atoi(s)
	if err != nil {panic(err)}
	fmt.Println(i)
}
