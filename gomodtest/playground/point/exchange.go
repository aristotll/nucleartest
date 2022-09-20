package main

import "fmt"

func exchangeTwoNum(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func main() {
	a, b := 10, 999
	fmt.Printf("交换前，a的值为 %d，b的值为 %d \n", a, b)
	exchangeTwoNum(&a, &b)
	fmt.Printf("交换后，a的值为 %d，b的值为 %d \n", a, b)
}
