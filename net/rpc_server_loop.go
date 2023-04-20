package main

import (
    "fmt"
    "net/rpc"
    "net"
    "math/rand"
    "time"
)

type Struct struct {}

type Req struct {}

type Resp struct {
    R []string
}

func (s *Struct) Do(req *Req, resp *Resp) error {
    var count int64
	var urls []string
	for {
		for count < 50 {
			urls = append(urls, genURL())
			count++
		}
        copy(urls, resp.R)	
		count = 0
		urls = urls[:0]
	}
}

func genURL() string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, 10)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return fmt.Sprintf("http://%s.com", string(b))
}

func main() {
    lis, err := net.Listen("tcp", ":9090")
    if err != nil {
        panic(err)
    }
    rpc.Register(&Struct{})
    for {
        conn, err := lis.Accept()
        if err != nil {
            fmt.Println(err)
            continue
        }
        go rpc.ServeConn(conn)
    }
}
