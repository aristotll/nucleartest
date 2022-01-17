package main

import (
	"os/exec"
	"log"
	"os"
)

func main() {
	log.Println("main")
	cmd := exec.Command("/proc/self/exe")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Start()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("end")	
}
