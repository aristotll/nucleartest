package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// 1.3 尝试测量可能低效的程序和使用 strings.Join 的程序在执行时间上的差异

// 低效程序：使用 + 拼接 string
func stringAppendTime() time.Duration {
	str := ""
	start := time.Now()
	for i := 0; i < 1000; i++ {
		str += strconv.Itoa(i) + strconv.Itoa(i*5)
	}
	end := time.Now()
	sub := end.Sub(start)
	return sub
}

// 使用 strings.Join 拼接
func joinTime() time.Duration {
	s := make([]string, 0)
	start := time.Now()
	for i := 0; i < 1000; i++ {
		s = append(s, strconv.Itoa(i) + strconv.Itoa(i*5))
		strings.Join(s, strconv.Itoa(i))
	}
	end := time.Now()
	sub := end.Sub(start)
	return sub
}

// 额外的测试: 使用新的 strings.Builder
func sbTime() time.Duration {
	var sb strings.Builder
	start := time.Now()
	for i := 0; i < 1000; i++ {
		sb.WriteString(strconv.Itoa(i) + strconv.Itoa(i*5))
	}
	end := time.Now()
	sub := end.Sub(start)
	return sub
}

func main() {
	// 8.698494ms
	t := joinTime()
	fmt.Println(t)

	// 1.399377ms
	t1 := stringAppendTime()
	fmt.Println(t1)

	// 94.301µs
	t3 := sbTime()
	fmt.Println(t3)
}
