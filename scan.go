package main

import (
    "bufio"
    "os"
    "fmt"
)

func main() {
    input := bufio.NewScanner(os.Stdin)
    input.Scan()
    text := input.Text()
    fmt.Println("输入的内容：", text)
}
