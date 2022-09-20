package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []string{"banana", "orange", "apple", "grapes"}
	sort.Strings(s)
	fmt.Println(s)

    s = []string{"day","is","sunny","the"}
    sort.Strings(s)
    fmt.Println(s)

    s = []string{"love","i"}
    sort.Strings(s)
    fmt.Println(s)

}
