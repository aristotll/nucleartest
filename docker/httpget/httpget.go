package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
)

var hostServiceEnvName = flag.String("hostEnv", "", "k8s service host env name")
var portServiceEnvName = flag.String("portEnv", "", "k8s service port env name")

func main() {
	flag.Parse()
	host := os.Getenv(*hostServiceEnvName)
	port := os.Getenv(*portServiceEnvName)
	url_ := fmt.Sprintf("http://%v:%v", host, port)
	fmt.Printf("%v:%v\n", *hostServiceEnvName, host)
	fmt.Printf("%v:%v\n", *portServiceEnvName, port)
	fmt.Printf("url: %v\n", url_)
	resp, err := http.Get(url_)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}
