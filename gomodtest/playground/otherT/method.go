package main

import (
	"fmt"
	"strings"
)

type M1 struct {

}

// 只能用指针类型接收
func (m *M1) String() string {
	return ""
}

func main() {
	m := new(M1)
	var m1 M1
	m.String()
	// 这里是编译器自动转换
	m1.String()

	lower := strings.ToLower("41D45CCCA459A3B423ACBD4F8E03A1CF")
	fmt.Println(lower)
}
