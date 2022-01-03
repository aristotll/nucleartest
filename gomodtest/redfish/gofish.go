package main

import (
    "fmt"

    "github.com/stmcginnis/gofish"
)

func main() {
    c, err := gofish.ConnectDefault("10.0.120.13")
    if err != nil {
        panic(err)
    }

    service := c.Service
    chassis, err := service.Chassis()
    if err != nil {
        panic(err)
    }

    for _, chass := range chassis {
        fmt.Printf("Chassis: %#v\n\n", chass)
    }
}
