package main

import "fmt"

type stu struct {
	name string
	age int
	score float64
}

func basicType(a *int) int {
	// a 中存储的是一个指针，所以虽然方法拷贝了一份参数，但里面的值仍然相同，可以通过*取出指针内容
	fmt.Printf("func param point is : %p, v is : %d \n", &a, a)
	fmt.Printf("before change *param is : %d \n", *a)
	*a = 500
	*a++
	fmt.Printf("change than *param is : %d, v is : %d \n", *a, a)
	return *a
}

func tBasicType(a *int) {
	r := basicType(a)
	fmt.Printf("tFunc param point is : %p, v is : %d \n", &a, a)
	fmt.Println(r)
}

func structType(s *stu) *stu {
	return s
}

func main() {
	a := 10
	fmt.Println("main before v:", a)
	tBasicType(&a)
	fmt.Printf("the point of main param is : %p, v : %d \n", &a, a)
}
