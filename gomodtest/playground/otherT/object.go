package main

import "fmt"

type person struct {
	name string
	age int
}

func (p *person) printPersonalInfo() {
	fmt.Println(p.age)
	fmt.Println(p.name)
}

func main() {
	p1 := &person{
		name: "zhang3",
		age:  22,
	}
	fmt.Printf("%T\n", p1)
	p1.name = "li4"

	p1.printPersonalInfo()

	var p2 *person = nil
	// 空指针异常
	if p2 != nil {
		p2.printPersonalInfo()
	}
}
