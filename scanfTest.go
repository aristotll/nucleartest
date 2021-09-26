
package main

import "fmt"

func main() {
    var s int
    fmt.Println("please input a number: ")
    fmt.Scanf("number: %d", &s)
    fmt.Println("s: ", s)
    
    fmt.Println("Please input a number again: ")
    fmt.Scanf("number: %d", &s)
    fmt.Println("2 s:", s)
}
