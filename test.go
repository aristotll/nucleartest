package main

import "fmt"
import "time"

func test(n int) (int, int, bool) {
	m := make(map[int]struct{})

	var n1, n2 int
	// 添加该范围内的所有素数
	for i := 2; i < n; i++ {
		if pre(i) {
			m[i] = struct{}{}
		}
	}

	for k := range m {
		// 能否整除
		if n%k == 0 {
			need := n / k
			// fmt.Printf("n: %d k: %d  need: %d \n", n, k, need)
			if _, ok := m[need]; ok {
				n1 = k
				n2 = need
				return n1, n2, true
			}
		}
	}

	return 0, 0, false
}

// 验证素数函数是否正确
func test1() {
	l := make([]int, 0)
	for i := 2; i < 100; i++ {
		if pre(i) {
			l = append(l, i)
		}
	}
	fmt.Println(l)
}

func test2(n int) (int, int, bool) {
	flag := false
	var n1, n2 int
	for i := 2; i <= n; i++ {
		if n%i != 0 {
			continue
		}
		if !pre(i) || !pre(n/i) {
			continue
		}
		flag = true
		n1 = i
		n2 = n / i
		break
	}
	return n1, n2, flag
}

func pre(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	t := time.Now()
	fmt.Printf("%s\n", "hello world")
	var p int
	fmt.Println("input a num:")
	fmt.Scanf("%d", &p)
	a, b, c := test2(p)
	fmt.Printf("%d %d %v\n", a, b, c)
	t1 := time.Since(t)
	fmt.Println("程序执行用时：", t1)
	//test1()
}
