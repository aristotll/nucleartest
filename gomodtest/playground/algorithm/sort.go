package main

import "fmt"

type Arr []float64

// 冒泡排序
func Bubble(arr Arr) Arr {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func main() {
	arr := make(Arr, 0)
	arr = append(arr, 123, 45, 6354, 23, 11)
	sArr := Bubble(arr)
	fmt.Println(sArr)
}
