package main

import (
	"fmt"
	"strconv"
)

type Employee struct {
	Id int
	Name string
	Salary float64
}

func GetEmployeeById(id int) *Employee {
	return &Employee{
		Name: strconv.Itoa(id) + "abc",
	}
}

func main() {
	employee := GetEmployeeById(12)
	employee.Salary = 500
	fmt.Println(employee)
}
