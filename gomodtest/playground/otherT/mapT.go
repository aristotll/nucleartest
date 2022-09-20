package main

import "fmt"

// key 是否存在
func isKeyExist(m *map[string]string) bool{
	v, ok := (*m)["a"]
	fmt.Println(v)
	return ok

}

func main() {
	m := make(map[string]string)
	m = map[string]string{
		"a": "123",
		"b": "456",
	}
	exist := isKeyExist(&m)
	fmt.Println("this key is ", exist)
}
