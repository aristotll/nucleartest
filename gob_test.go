package main

import (
	"bufio"
	"encoding/gob"
	"net"
	"testing"
)

type A struct {
	X int64
}

type B struct {
	XX int64
	YY string
}

const addr = ":8999"

func TestSocketServer(t *testing.T) {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		t.Fatal(err)
	}
	for {
		conn, err := lis.Accept()
		if err != nil {
			t.Log(err)
			continue
		}
		t.Logf("new client: %v", conn.RemoteAddr())
		go func() {
			var a A
			var b B
			for {
				// 必须用 bufio 包装，否则解码可能会出现这些错误情况： EOF、
				// gob: unknown type id or corrupted data 以及无限阻塞，但有时能解析成功
				br := bufio.NewReader(conn)
				t.Log("is gob decode block?") // 不会 block
				// 客户端是先编码 a 再编码 b，服务端这里先解码 b 再解码 a 会怎样？
				// 如果是不同的结构体则直接报错，相同的结构体会导致 b 的数据解析为 a，a 的解析为 b
				if err := gob.NewDecoder(br).Decode(&a); err != nil {
					t.Log(err)
				}
				t.Log("test block")
				if err := gob.NewDecoder(br).Decode(&b); err != nil {
					t.Log(err)
				}
			}
			conn.Close()
		}()
	}
}

func TestSocketClient(t *testing.T) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		t.Fatal(err)
	}
	a := &A{1}
	b := &B{2, "abc"}
	if err := gob.NewEncoder(conn).Encode(&a); err != nil {
		t.Fatal(err)
	}
	if err := gob.NewEncoder(conn).Encode(&b); err != nil {
		t.Fatal(err)
	}
}

// 使用 nil 解码会怎样？
func TestSocketServer1(t *testing.T) {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		t.Fatal(err)
	}

	for {
		conn, err := lis.Accept()
		if err != nil {
			t.Log(err)
			continue
		}
		br := bufio.NewReader(conn)
		// 使用 nil 解码会怎样？
		if err := gob.NewDecoder(br).Decode(nil); err != nil {
			t.Log(err)
			return
		}
		// 用 nil 解码后，conn 中还有数据吗？
		// 结果：无数据
		//var a A
		//if err := gob.NewDecoder(br).Decode(&a); err != nil {
		//	t.Log(err)
		//	return
		//}
		//t.Log(a)
		var b B
		if err := gob.NewDecoder(br).Decode(&b); err != nil {
			t.Log(err)
			return
		}
		t.Log(b)
	}
}

func TestSocketClient1(t *testing.T) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		t.Fatal(err)
	}
	a := &A{1}
	b := &B{2, "abc"}
	if err := gob.NewEncoder(conn).Encode(&a); err != nil {
		t.Fatal(err)
	}
	if err := gob.NewEncoder(conn).Encode(&b); err != nil {
		t.Fatal(err)
	}
}
