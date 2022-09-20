package main

import (
	"fmt"
	"os"
	"runtime/pprof"
	_ "runtime/pprof"
)

func fibFast(n int) int {
	if n == 0 {
		return 0
	}
	if n <= 2 {
		return 1
	}

	var a, b = 1, 1
	var c int

	for i := 3; i <= n; i++ {
		c = a + b
		a = b
		b = c
	}

	return c
}

func fibSlow(n int) int {
	if n <= 1 {
		return 1
	}

	return fibSlow(n-1) + fibSlow(n-2)
}

func main() {
	f, _ := os.OpenFile("cpu.profile", os.O_CREATE|os.O_RDWR, 0777)
	defer f.Close()
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	n := 10
	for i := 1; i <= 5; i++ {
		fmt.Println(fibSlow(n))
		n += 3 * i
	}
}
