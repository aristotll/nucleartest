package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.OpenFile("peek_test", os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	ws1 := "123"
	ws2 := "456"
	w.WriteString(ws1)
	w.WriteString(ws2)
	w.Flush()

	f.Seek(0, 0)

	r := bufio.NewReader(f)
	// peek 查看指定长度的内容，不会移动指针，会返回查看的内容
	// 这里 peek 的长度为 3，这正好是 ws1 的长度
	b, err := r.Peek(len(ws1))
	if err != nil {
		panic(err)
	}
	// Output: 123
	// 输出的是 ws1，符合预期
	fmt.Println(string(b))

	// 多次 peek 结果相同
	b, err = r.Peek(len(ws1))
	if err != nil {
		panic(err)
	}
	// Output: 123
	fmt.Println(string(b))

	// 验证 peek 不会移动指针
	// buf := make([]byte, 10)
	// _, err = r.Read(buf)
	// if err != nil {
	// 	panic(err)
	// }
	// // Output: 123456
	// // 这里依然输出了文件的全部内容，没有受到 peek 的影响
	// fmt.Println(string(buf))

	// 如果想实现移动指针的效果，可以使用 Discard 来跳过指定长度的内容
	r.Discard(len(b))
	b, err = r.Peek(len(ws2))
	if err != nil {
		panic(err)
	}
    // Output: 456
	fmt.Println("after discard: ", string(b))
    copy()

	// 如果 peek 的长度超过了读出的内容的长度，则会 panic
	b, err = r.Peek(100)
	if err != nil {
		panic(err)
	}
}
