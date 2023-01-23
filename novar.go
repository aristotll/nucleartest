package main

import (
    "fmt"
)

type S struct {
    labels map[string]string
}

func (s *S) setLabels(labels map[string]string) {
    s.labels = labels
}

func main() {
    s := new(S)
    s.labels["1"] = "1"
    s.setLabels(make(map[string]string))
    s.labels["1"] = "1"
    fmt.Println(s)
}

