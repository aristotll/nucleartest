package main

import (
	"time"

	"github.com/miekg/dns"
)

func main() {
	c := dns.Client{
		Timeout: time.Second * 5,
	}
	conn, err := c.Dial("")
	if

	var m dns.Msg

}