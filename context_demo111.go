package main

import (
	"context"
	"log"
	"net/http"
	"os/exec"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*4)
	defer cancel()

	exec.CommandContext(ctx, "sleep", "3").CombinedOutput()
	log.Println("ok")

	_, err := exec.CommandContext(ctx, "sleep", "3").CombinedOutput()
	if err != nil {
		log.Println(err)
	}
	log.Println("ok")

	ctx1, cancel := context.WithTimeout(context.Background(), time.Hour)
	r, err := http.NewRequestWithContext(
		ctx1,
		"get",
		"https://www.google.com",
		nil)

	_, err = http.DefaultClient.Do(r)
	if err != nil {
		log.Fatalln(err)
	}
}
