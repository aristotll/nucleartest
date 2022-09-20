package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

// 缓冲区读取文件

func main() {
	file, err := os.OpenFile("io/t.txt", os.O_RDONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	var b [50]byte
	var result string

	reader := bufio.NewReader(file)

	for {
		n, err := reader.Read(b[:])
		if err != io.EOF {
			result += string(b[:n])
			// fmt.Println(string(b[:n]))
		}else {
			break
		}
	}
	fmt.Println(result)



	//for {
	//	line, _, err := reader.ReadLine()
	//	if err != io.EOF {
	//		fmt.Println(string(line))
	//	}else {
	//		break
	//	}
	//}

}
