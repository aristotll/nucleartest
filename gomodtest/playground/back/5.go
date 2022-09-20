package main

import "fmt"

func del(nums [5]int) {
	nums[0] = 1231
}

func delP(nums *[5]int) {
	(*nums)[0] = 1231
}

func delM(m map[string]string) {
	m["abc"] = "123"
}


func main() {
	n := [...]int{1, 2, 3, 4, 5}
	del(n)
	fmt.Println(n)

	delP(&n)
	fmt.Println(n)

	m := make(map[string]string)
	m["abc"] = "abc"
	delM(m)
	fmt.Println(m)
}
