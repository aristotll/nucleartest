package main

import (
	"bufio"
	"fmt"
	"os"
)

func dup() {
	m := make(map[string]int)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		m[scanner.Text()]++
		if scanner.Text() == "q" {
			fmt.Println("退出")
			break
		}
	}
	for k, v := range m {
		if v > 1 {
			fmt.Println("重复行 ========>", k, v)
		}
	}
}

func main() {
	dup()
}
