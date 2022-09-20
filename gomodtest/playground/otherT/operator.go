package main

import "fmt"

func main() {
	fmt.Printf("3 = %b 5 = %b \n", 3, 5)
	fmt.Printf("%b \n", 3&5)

	fmt.Printf("2 = %b 4 = %b \n", 2, 4)
	fmt.Printf("base 2: %b base 10: %d\n", 2|4, 2|4)

	fmt.Println(4 >> 1)
	fmt.Printf("1 << 1 = %d\t1 << 2 = %d\t1 << 3= %d\t1 << 5 = %d\n",
		1<<1, 1<<2, 1<<3, 1<<5)

	fmt.Println(1<<1 | 1<<5)
	fmt.Println(1<<1 | 1<<2)


}
