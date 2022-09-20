package private

import (
	"fmt"
	_ "unsafe"
)

// 使用 go:linkname 将该函数链接到 pub 包下的 callPrivateFunc 函数
//go:linkname privateFunc golinkname/pub.callPrivateFunc
func privateFunc(name string) {
	fmt.Printf("[you call a private func] name:%s \n", name)
}

func privateFunc1(name string, age int8) {
	fmt.Printf("[you call a private func] name:%s, age:%d \n", name, age)
}
