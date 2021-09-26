package main

import (
	"fmt"
	"github.com/huandu/skiplist"
)

func main() {
	l := skiplist.New(skiplist.Int)
	l.Set(11, "aa")
	l.Set(11, "bb")

	v := l.Get(11)
	fmt.Println(v.Value)
}
