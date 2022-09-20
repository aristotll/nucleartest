package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func read() {
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("输入：")
		text, err := reader.ReadString('\n')
		if err != nil{
			return
		}
		text = strings.TrimSpace(text)
		if text == "q" {
			return
		} else if text == "foo"{
			fmt.Println("foo")
		}else{
			fmt.Println("bar")
		}
	}
}

func array() {
	var arr []int
	arr = append(arr, 123)
	arr = append(arr, 456)
	for _, v := range arr {
		fmt.Println(v)
	}
}

func main() {
	// read()
	array()
	fmt.Println(7 >> 2)
}
