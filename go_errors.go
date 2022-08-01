package main

import (
	"errors"
	"fmt"
)

func dao() error {
	err1 := errors.New("dao err1")
	err2 := fmt.Errorf("dao err2: %w", err1)
	return fmt.Errorf("dao err3: %w", err2)
}

func service() error {
	err := dao()
	return fmt.Errorf("service error: %w", err)
}

func controller() {
	err := service()
	fmt.Println(err)
}

func main() {
	controller()
}
