package main

import (
	"os/exec"
	"log"
)

func main() {
	cmd := exec.Command("bash", "-c", "ls -l")
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(output))
}
