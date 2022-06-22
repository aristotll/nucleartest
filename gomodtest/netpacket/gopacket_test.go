package netpacket

import (
	"fmt"
	"log"
	"testing"
)

func init() {
	log.SetFlags(log.Lshortfile | log.Ldate)
}

func TestSearchDevice(t *testing.T) {
	SearchDevice()
}

func TestSendPacket(t *testing.T) {
	SendPacket()
}

func TestSendRST(t *testing.T) {
	seq := SendSYN()
	SendRST(seq)
}

func TestSendSYN(t *testing.T) {
	fmt.Printf("SendSYN(): %v\n", SendSYN())
}
