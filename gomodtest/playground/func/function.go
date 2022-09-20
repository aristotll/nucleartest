package main

import "fmt"

var sum = func(f func(a, b int) int, num1, num2 int)  {
	result := f(num1, num2)
	fmt.Println("test", result)
}

var forEach = func(f func(o []interface{}), arr []interface{}) {
	f(arr)
}


func main() {
	sum(func (a, b int) int {
		return a + b
	}, 1, 3)

	arr := make([]interface{}, 10)
	strings := []string{"aa", "bb", "dd"}
	for k, v := range strings {
		arr[k] = v
	}
	i := make(map[string]string)

	i["country"] = "巴黎"



	forEach(func(i []interface{}) {
		fmt.Println("result", i)
	}, arr)
	
}
