package main

import (
	"fmt"
	"sort"
)

func main() {
	var s []string = []string{"ab", "hg", "cc", "巴巴爸爸", "啊啊啊"}
	sort.Strings(s)
	
	fmt.Println(s)
}
