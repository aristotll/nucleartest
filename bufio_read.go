package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	i := bufio.NewReader(os.Stdin)
	fmt.Printf("Please enter some input: ")
	input, err := i.ReadString('\n')
	if err == nil {
		fmt.Printf("The input was: %s", input)
	}

	input, _ = i.ReadString('\n')
	if err == nil {
		fmt.Printf("input: %s", input)
	}
}
