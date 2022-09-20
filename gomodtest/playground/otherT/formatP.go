package main

import "fmt"

func main() {
	s := make([]string, 0)
	s = append(s, "123")
	s = append(s, "发达")
	s = append(s, "wtt")

	str := "123"
	// ["123" "发达" "wtt"] [123 发达 wtt]
	fmt.Printf("%q %v \n", s, s)
	// s = 123 q = "123"
	fmt.Printf("s = %s q = %q \n", str, str)

	sp := fmt.Sprintf("%q", s)
	fmt.Println(sp)
}
