package main

import (
	"fmt"
	"os"
	"os/signal"
)

func main() {
	signch := make(chan os.Signal, 1)
	signal.Notify(signch, os.Interrupt)
	<-signch
	fmt.Println("SIG INTER")
}
