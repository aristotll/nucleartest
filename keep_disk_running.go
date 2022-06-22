package main

import (
	"log"
	"os"
	"time"
)

func main() {
	p := "/Volumes/4t/DO_NOT_SLEEP"
	for {
		f, err := os.OpenFile(p, os.O_CREATE|os.O_RDWR, 0777)
		if err != nil {
			log.Println(err)
			time.Sleep(time.Second * 5)
			continue
		}
		f.WriteString("123")
		log.Println("write ok")
		f.Close()
		time.Sleep(time.Second * 5)
	}
}
