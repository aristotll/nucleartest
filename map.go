package main

import (
    //"fmt"
)

func main() {
    var m map[string]string
    _ = m["1"]
    //m["1"] = "1"
    delete(m, "1")
}
