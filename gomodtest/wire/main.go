package main

import (
	"fmt"
	"log"
	"wire/wirefile"
)

// 使用 wire 命令生成文件
// wire gen [文件路径]
func main() {
	// 调用该方法会返回一个 Baz 对象，其所依赖的对象都已经注入完成
	baz, err := wirefile.InitializeBaz(123)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(baz.X)

	bazz, err := wirefile.InitializeBazz(456)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(bazz)
	
}


