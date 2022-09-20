package main

import (
	"log"
	"os"
	"time"
)

func main() {
	p := "/Volumes/4t/DO_NOT_SLEEP"
    var recorded bool
	for {
		f, err := os.OpenFile(p, os.O_CREATE|os.O_RDWR, 0777)
		if err != nil {
            if recorded == false {
			    log.Println(err)
                recorded = true
            }
			time.Sleep(time.Second * 10)
			continue
		}
		f.WriteString("\n")
		//log.Println("write ok")
		f.Close()
		time.Sleep(time.Second * 10)
	}
}
