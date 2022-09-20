package main

import (
	"fmt"
)

func checkRoles(roles []string) bool {
	var a, u bool
	for _, v := range roles {
		if v == "admin" {
			a = true
		} else if v == "user" {
			u = true
		}
	}
	return a && u
}

func main() {
	s := []string{"aaa", "admin"}
	fmt.Println(checkRoles(s))

	s = []string{"user", "admin"}
	fmt.Println(checkRoles(s))

	s = []string{"aa", "bb"}
	fmt.Println(checkRoles(s))
}
