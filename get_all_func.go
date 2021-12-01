package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

var path = flag.String("p", "", "input file path")

func main() {
	flag.Parse()

	f, err := os.Open(*path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	b, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}
	content := string(b)

	for len(content) > 0 {
		i := strings.Index(content, "func") // 找到第一个 func
		i2 := strings.Index(content[i:], "(")   // 找到第一个 func 之后的第一个 (
		if i2 == -1 {
			break
		}
		funcName := strings.TrimSpace(content[i+4 : i2+i]) // func 和 (  之间的就是函数名
		fmt.Println(funcName)

		i = strings.Index(content[i2+i:], "func")
		content = content[i:] // 截取字符串，从第二个 func 开始
		//fmt.Println(content)
	}

}
