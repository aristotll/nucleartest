package main

import "fmt"

type stu struct {
	name string
	age int
}

// 结构体指针
func main() {
	s := new(stu)
	s2 := &stu{
		name: "zzz",
		age:  16,
	}
	// &s: 0xc00000e028 s: 0xc00000c060
	fmt.Printf("&s: %p s: %p \n", &s, s)
	// &s2: 0xc00000e030, s2: 0xc00000c080
	fmt.Printf("&s2: %p, s2: %p \n", &s2, s2)

	// 将 s2 的地址赋给 s
	s = s2
	// s: 0xc00000c080, &s: 0xc00000e028, *s: {zzz 16}
	fmt.Printf("s: %p, &s: %p, *s: %v \n", s, &s, *s)
	// 如下，直接输出 s 和 *s，结果是类似的，只是 s 的结果多了一个取址符
	// 直接输出结构体指针，结果不是地址，而是值
	// &{123 112}
	fmt.Println(s)
	// {123 112}
	fmt.Println(*s)
}
