package main

import "fmt"

func sumSequence(n int) float64 {
	if n < 1 {
		panic("error input!")
	}

	var sum = 0.0
	var startSon = 1.0
	var startMother = 2.0

	for i := 0; i <= n; i++ {
		sum += startSon / startMother
		var newMother = startSon + startMother
		startSon = startMother
		startMother = newMother
		onceResult := startSon / startMother
		fmt.Printf("r2 : %f / %f = %f \n",
			startSon, startMother, onceResult)
	}
	return sum
}

func main() {
	fmt.Println("please input number: ")
	var in int
	_, _ = fmt.Scanf("%d", &in)
	result := sumSequence(in)
	fmt.Println("最终结果：", result)
}
