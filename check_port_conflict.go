package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"net/http"
	"sync"
	"time"
)

func GetRandomPort() int32 {
	rand.Seed(time.Now().Unix())
	for {
		port := rand.Intn(10000) + 40000
		_, err := net.Dial("tcp", fmt.Sprintf("localhost:%v", port))
		if err != nil {
			continue
		} else {
			return int32(port)
		}
	}
}

func startHTTPServer(port int32) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

	})
	addr := fmt.Sprintf("localhost:%v", port)
	log.Println("http start listen ")
	if err := http.ListenAndServe(addr, nil); err != nil {
		panic(err)
	}
	
}

func checkPortConflict(port int32) (conflict bool) {
	_, err := net.Listen("tcp", fmt.Sprintf("localhost:%v", port))
	if err != nil {
		return true
	}
	return false
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		startHTTPServer(8080)
	}()
	time.Sleep(time.Second * 2)
	fmt.Println(checkPortConflict(8080))
	wg.Wait()
}
