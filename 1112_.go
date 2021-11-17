package main

import (
    "fmt"
)

func distinctRoles(roles []string) []string {
	m := make(map[string]struct{})
	for _, v := range roles {
		m[v] = struct{}{}
	}

	var s []string
	for k := range m {
		s = append(s, k)
	}

	return s
}

func main() {
    s := []string{"admin", "admin", "user", "user"}
    s = distinctRoles(s)
    fmt.Println(s)

    s1 := []string{"admin", "admin"}
    s1 = distinctRoles(s1)
    fmt.Println(s1)
}
