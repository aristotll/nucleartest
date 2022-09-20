package main

import (
	"flag"
	"fmt"
	"net"
)

var url = flag.String("u", "", "input a url")

func main() {
	flag.Parse()
	//fmt.Println(*url)
	ns, err := net.LookupHost(*url)
	if err != nil {
		fmt.Printf("error: %v, failed to parse %v\n", err, *url)
		return
	}

	fmt.Printf("parse %v \n", *url)

	for _, i := range ns {
		fmt.Printf("%s\n", i)
	}
}
