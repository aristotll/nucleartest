package main

import (
	"fmt"
)

func backtrack(s, cur string, res *[]string, start int) {
	if start == len(s) {
		*res = append(*res, cur)
		return
	}

	b := []byte(cur)

	for i := start; i < len(s); i++ {
		b[i], b[start] = b[start], b[i]
		backtrack(s, string(b), res, start+1)
		b[i], b[start] = b[start], b[i]
	}

}

func main() {
	var res []string
	backtrack("qwe", "qwe", &res, 0)
	fmt.Println(res)
}
