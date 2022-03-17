package main

import (
	"sort"
	"fmt"
)

func main() {
	s := []string{"banana","orange","apple","grapes"}
	sort.Strings(s)
	fmt.Println(s)
}
