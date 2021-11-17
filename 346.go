package main

import (
    "strings"
    "fmt"
    "unicode"
)

func main() {
    s := "my name is a, and her name is b."
    ff := func(r rune) bool {return !unicode.IsLetter(r)}

    word := strings.FieldsFunc(s, ff)

    fmt.Println(word)
}
