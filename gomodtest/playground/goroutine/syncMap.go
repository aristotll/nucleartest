package main

import (
	"fmt"
	"sync"
)

// 并发下 map 不安全的例子

var m = make(map[int]int)

type p struct {
	age int
	name string
}

// fatal error: concurrent map read and map write
// 错误信息显示，并发的 map 读和 map 写，
// 也就是说使用了两个并发函数不断地对 map 进行读和写而发生了竞态问题，
// map 内部会对这种并发操作进行检查并提前发现
func MapCurrent() {
	go func() {
		for {
			m[1] = 1
		}
	}()

	go func() {
		for {
			_ = m[1]
		}
	}()

	for {
		_ = m[1]
	}
}

func SafeMap() {
	// 直接声明即可
	var syncM sync.Map

	// Store(): 存储键值对
	syncM.Store("fku", 123)
	syncM.Store(&p{
		age:  12,
		name: "wang",
	}, "123")
	syncM.Store(0x142F, "kkkk")

	load, ok := syncM.Load("fku")
	if ok {
		fmt.Println(load)
	}

	syncM.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		return true
	})
}

func main() {
	// MapCurrent()
	SafeMap()
}
