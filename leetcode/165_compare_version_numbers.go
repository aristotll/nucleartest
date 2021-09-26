package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	//"bytes"
)

// v1 > v2 => 1, v1 < v2 => -1
func compare(v1, v2 string) int {
	start := time.Now()
	v1arr := strings.Split(v1, ".")
	v2arr := strings.Split(v2, ".")
	//
	if len(v1arr) > len(v2arr) {
		sub := len(v1arr) - len(v2arr)
		for sub > 0 {
			v2arr = append(v2arr, "0")
			sub--
		}
	} else if len(v1arr) < len(v2arr) {
		sub := len(v2arr) - len(v1arr)
		for sub > 0 {
			v1arr = append(v1arr, "0")
			sub--
		}
	}
	fmt.Printf("v1: %v v2:%v\n", v1arr, v2arr)

	for i := 0; i < len(v1arr); i++ {
		v1i, _ := strconv.Atoi(v1arr[i])
		v2i, _ := strconv.Atoi(v2arr[i])
		if v1i > v2i {
			return 1
		} else if v1i < v2i {
			return -1
		} else {
			continue
		}
	}

	//s, _ := strconv.Atoi("001")
	//fmt.Println(s)
	runTime := time.Since(start)
	fmt.Println("running time: ", runTime)
	return 0
}

// 双指针
func compareTwoPoint(v1, v2 string) int {
	p1, p2 := 0, 0
	l1, l2 := len(v1), len(v2)

	var maxFunc = func(x, y int) int {
		if x > y {
			return x
		} else {
			return y
		}
	}

	max := maxFunc(l1, l2)
	//vv1, vv2 := 0, 0

	for p1 < max || p2 < max {
        vv1, vv2 := 0, 0

        // 使用 while 来循环读取一个小版本号（例如 1.111 中的 1 和 111 就是小版本号），
        // 遇到 . 停止，此时 vv1, vv2 的值即是小版本号，vv1 和 vv2 定义在外层 for 内，
        // 每次比较后都会清零
		for p1 < l1 && v1[p1] != '.' {
			vv1 = vv1*10 + int(v1[p1]) - '0'
            p1++
		}
		for p2 < l2 && v2[p2] != '.' {
			vv2 = vv2*10 + int(v2[p2]) - '0'
            p2++
		}
		//fmt.Printf("vv1: %d vv2: %d\n", vv1, vv2)
		if vv1 > vv2 {
			return 1
		} else if vv2 > vv1 {
			return -1
		}
		p1++
		p2++
	}
	return 0
}

func main() {
	//v1 := "1.12.0.1"
	//v2 := "1.12.0.001"
	//v1 := "7.5.2.4"
	//v2 := "7.5.3"
	//v1 := "1.0.1"
	//v2 := "1"
	v1 := "1.0"
	v2 := "1.0.0"
	//v1 := "1.01"
	//v2 := "1.001"
	//r := compare(v1, v2)
	r := compareTwoPoint(v1, v2)
	fmt.Println(r)
}
