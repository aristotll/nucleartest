package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

type W func(p string) []string

// 读取指定路径下的全部文件名，写入到一个 slice 中
// @param path 路径
func writeList(path string) W {
	return func(p string) []string {
		p = path
		var list []string
		dir, err := ioutil.ReadDir(p)
		if err != nil {
			fmt.Println("读取错误: ", err)
		}

		for _, v := range dir {
			list = append(list, v.Name())
			// fmt.Println(v.Name())
			// fmt.Println(ll)
		}
		return list
	}
}

// demo: 获取指定目录下的所有文件名，并写入到指定目录下的一个 txt 中
// @param writePath: txt的写入路径
// @param readPath:  要读取的目录路径
// @param f: 读取指定路径下的全部文件名，写入到一个 slice 中
//           是一个 func(p string) []string
func handle(writePath string, f W) {
	// list := f(readPath)
	list := f("")
	file, err := os.OpenFile(writePath + "导出.txt",
		os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("创建 txt 错误, err:", err)
	}
	for _, v := range list {
		_, err := file.Write([]byte(v + "\n"))
		if err != nil {
			fmt.Println("将文件名批量写入txt失败，错误：", err)
		}
		// fmt.Println(v)
	}
}

func main() {
	fmt.Println("输入要读取的路径")
	var readPath string
	_, err := fmt.Scanf("%s", &readPath)
	if err != nil {
		fmt.Println("控制台读取错误, err:", err)
	}

	fmt.Println("输入要写入的路径, 记得末尾加/")
	var writePath = ""
	_, _ = fmt.Scanf("%s", &writePath)
	// /Users/zeng/Downloads
	fmt.Println("write path: ", writePath)

	handle(writePath, writeList(readPath))
}
