package main 

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("os.Args[0] = %v \n", os.Args[0])
	fmt.Printf("os.Args[0] == /proc/self/exe ? %v \n", os.Args[0] == "/proc/self/exe")
}
