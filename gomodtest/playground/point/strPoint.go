package main

import "fmt"

type names struct {
	name *string
	age *int
}

func strT(s *string) names{
	fmt.Println(*s)
	st := "asdasda"
	i := 123

	n := names{
		name: &st,
		age:  &i,
	}
	return n
}

func t11() {
	i := 123
	a := &i
	fmt.Println(a)
	fmt.Println(*a)
}

func main() {
	str := new(string)
	// str = "123"
	*str = "123"

	n := strT(str)
	fmt.Printf("%s %d \n", *n.name, *n.age)

	t11()
}
