package main

import "fmt"

// set 通过 map的 key 唯一原则实现

// 该结构体用于作为 map 的 value，因为空结构体不占用内存
type null struct{}

type set struct {
	se map[interface{}]null
}

type stu struct {
	name string
	age int
}

func NewSet() *set {
	m := make(map[interface{}]null)
	return &set{se: m}
}

func (s *set) Add(v interface{}) *set {
	m := s.se
	// 如果该 key 不存在
	if _, ok := m[v]; !ok {
		s.se[v] = null{}
	}
	return s
}

func (s *set) Foreach() {
	for k, _ := range s.se {
		fmt.Print(k, "  ")
		// fmt.Printf("%d  ", k)
	}
}

func main() {
	// int 类型，测试通过
	NewSet().Add(5).Add(44).Add(446).Foreach()
	// string 测试通过
	NewSet().Add("asdd").Add("abcd").Add("abcd").Foreach()
	// struct 测试通过，相同值的结构体会自动过滤？
	NewSet().Add(stu{
		name: "zhang",
		age:  16,
	}).Add(stu{
		name: "zhang",
		age:  16,
	}).Foreach()

	s := stu{
		name: "zhang",
		age:  16,
	}

	s2 := stu{
		name: "zhang",
		age:  16,
	}

	if s == s2 {
		fmt.Println("true")
	}
	fmt.Printf("s: %p s2: %p", &s, &s2)
}
