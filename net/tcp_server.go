package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"math/rand"
	"net"
	"os"
	"strconv"
	"sync"
	"time"
)

type Client struct {
	Conn      net.Conn
	WriteChan chan []byte
	ReadChan  chan []byte
}

var wg sync.WaitGroup

func (c *Client) Read() {
	defer wg.Done()
	for {
		data, err := read(c.Conn)
		if err != nil {
			log.Println("server read message error: ", err)
			break
		}
		c.ReadChan <- data
	}
}

func read(conn net.Conn) ([]byte, error) {
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		return buf, err
		//fmt.Println("read error: ", err)
	}
	return buf[:n], nil
}

func (c *Client) Write(msg []byte) {
	defer wg.Done()
	n, err := write(c.Conn, msg)
	if err != nil {
		log.Println("server write message error: ", err)
	}
	log.Printf("read [%v]bytes\n", n)
	c.WriteChan <- msg

}

func write(conn net.Conn, msg []byte) (int, error) {
	// buf := make([]byte, 1024)
	rand.Seed(time.Now().Unix())
	intn := rand.Intn(math.MaxInt32)
	content := "rand id: " + strconv.Itoa(intn)
	v := content + string(msg)
	n, err := conn.Write([]byte(v))
	if err != nil {
		return -1, err
	}
	return n, nil
}

func (c *Client) handler() {
	defer wg.Done()
	defer c.Conn.Close()
	for {
		select {
		case msg := <-c.ReadChan:
			fmt.Println(string(msg))
		case msg := <-c.WriteChan:
			c.Conn.Write(msg)
		}
	}
	// wg.Add(2)
	// go write(c.Conn)
	// go read(c.Conn)
	// wg.Wait()
}

func main() {
	listener, err := net.Listen("tcp", ":8083")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("accept conn error: ", err)
			continue
		}
		cli := &Client{
			Conn:      conn,
			WriteChan: make(chan []byte),
			ReadChan:  make(chan []byte),
		}
		wg.Add(3)
		go cli.handler()
		go cli.Read()
		reader := bufio.NewReader(os.Stdin)
		data, _ := reader.ReadString('\n')
		go cli.Write([]byte(data))
		wg.Wait()
	}
}
