package main

import "fmt"

// 嵌套指针示例
func main() {
	i := 1545
	// i 是 int 类型，所以 i 的 指针类型为 *int
	var ii *int = &i
	// ii 是 *int 类型，所以 ii 的 指针类型为 **int
	var iii **int = &ii
	// iii 是 int 类型，所以 i 的 指针类型为 *i
	var iiii ***int = &iii

	// ii 存放的是 i 的地址，*ii 即 取出 i 的值
	fmt.Printf("%d \n", *ii)
	// iii 存放的是 ii 的地址，*iii 取出的是 ii 的值（i 的地址），所以*(*iii) = *(&i) = 1545
	// *iii = ii = &i => *(*iii) = *(&i) = 1545
	fmt.Printf("%d \n", *(*iii))
	// iiii 存放的是 iii 的地址，*iiii 取出的是 iii 的值，即 *iiii = iii
	// **(*iiii) = *(*iii) = 1545
	fmt.Printf("%d \n", **(*iiii))
}