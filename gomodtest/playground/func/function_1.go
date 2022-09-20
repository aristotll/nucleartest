package main

import "fmt"

var plus = func(a ...int) int{
	sum := 0
	for _, v := range a {
		sum += v
	}
	return sum
}

func cal(f func(num []float64) float64, num ...float64) float64 {
	var arr []float64
	for _, v := range num {
		arr = append(arr, v)
	}
	r := f(arr)
	return r
}

// var a = 0x3BEH
// var aa = 0x3fed   value:16365

func main() {
	//fmt.Printf("t: %T \n", aa)
	//fmt.Println(aa)
	r := plus(1, 2, 3)
	fmt.Println(r)

	// 加法
	add := cal(func(num []float64) float64 {
		sum := 0.0
		for _, v := range num {
			sum += v
		}
		return sum
	}, 1.0, 2.0, 3.0)
	fmt.Println("加法：", add)

	// 减法
	sub := cal(func(num []float64) float64 {
		firstNum := num[0]
		for i := 1; i < len(num); i++ {
			firstNum -= num[i]
		}
		return firstNum
	}, 10, 2, 3)
	fmt.Println("减法", sub)


}