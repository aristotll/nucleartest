package main

import (
	"errors"
	"log"
	"net"
	"strconv"
	//"strings"
	//"bufio"
)

// 一个网络计算器 demo，客户端发送 n 个数字和一个运算符，服务端在计算后将结果返回给客户端

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("create server error: ", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("accept error: ", err)
			continue
		}
		go func() {
			// 客户端可以持续输入
			for {
				// 接收客户端输入
				buf := make([]byte, 1024)
				_, err = conn.Read(buf)
				if err != nil {
					log.Println("read message error: ", err)
					break
				}

				// 输入 q 退出
				if buf[0] == 'q' || buf[0] == 'Q' {
				   conn.Write([]byte("^^ bye ^^"))
				   conn.Close()
				   break
				}

				op, num, err := ParseProtocol(buf)
				if err != nil {
					log.Println(err)
					continue
				}
				log.Printf("[parse] op: %v num: %v\n", op, num)

				res, err := handler(op, num)
				if err != nil {
					log.Println(err)
					continue
				}

				sres := strconv.Itoa(res)

				conn.Write([]byte(sres))
			}
		}()


	}
}

// ParseProtocol 解析协议信息
func ParseProtocol(msg []byte) (op byte, num []int, err error) {
	if len(msg) == 0 {
		return ' ', nil, errors.New("protocol info error")
	}
	count := int(msg[0]-'0') // 运算数个数

	// 读取运算数，+1 跳过[0]
	for i := 0; i < count; i++ {
		num = append(num, int(msg[i+1]-'0'))
	}

	op = msg[count+1] // 运算符
	return
}

func handler(op byte, num []int) (result int, err error) {
	if op != '+' && op != '-' && op != '*' && op != '/' {
		err = errors.New("Illegal operator")
	}

	result = num[0]
	for i := 1; i < len(num); i++ {
		if op == '+' {
			result += num[i]
		} else if op == '-' {
			result -= num[i]
		} else if op == '*' {
			result *= num[i]
		} else if op == '/' {
			result /= num[i]
		}
	}

	return
}
