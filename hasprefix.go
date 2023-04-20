package main

import (
	"fmt"
	"strings"
)

func main() {
	s := `---
k:v
kk:vv
---

jisajdiljifdlj
asjlkjslqwe
dfgdfgdf
asdase5er`

	b := strings.HasPrefix(s, "---")
	fmt.Println(b)

    s1 := `jisajdiljifdlj
asjlkjslqwe
dfgdfgdf
asdase5er`

    b = strings.HasPrefix(s1, "---")
	fmt.Println(b)

    fmt.Println(strings.Index(s, "---"))
}
