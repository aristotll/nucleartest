package main

import (
	"fmt"
	"strings"
	"syscall"

	"github.com/kr/pretty"
)

type NetlinkListener struct {
	fd int
	sa *syscall.SockaddrNetlink
}

func (l *NetlinkListener) ReadMsgs() (string, error) {
	defer func() {
		recover()
	}()

	pkt := make([]byte, 2048)
	_, err := syscall.Read(l.fd, pkt)
	if err != nil {
		return "null", fmt.Errorf("read: %s", err)
	}

	fmt.Println(pretty.Sprint(string(pkt)))

	outMsg := string(pkt)
	if find := strings.Contains(outMsg, "DEVTYPE=partition") && strings.Contains(outMsg, "SUBSYSTEM=block"); find {
		action := strings.Split(outMsg, "@")[0]
		tmp := strings.Split(outMsg, "ACTION")[0]
		name := strings.Split(tmp, "/")[len(strings.Split(tmp, "/"))-1]
		fmt.Println("ACTION:", action, ";", "DEVICE:", name)
	}
	return string(pkt), nil
}

func ListenNetlink() (*NetlinkListener, error) {
	groups := syscall.RTNLGRP_LINK |
		syscall.RTNLGRP_IPV4_IFADDR |
		syscall.RTNLGRP_IPV6_IFADDR

	s, err := syscall.Socket(syscall.AF_NETLINK, syscall.SOCK_DGRAM,
		syscall.NETLINK_KOBJECT_UEVENT)
	if err != nil {
		return nil, fmt.Errorf("socket: %s", err)
	}

	saddr := &syscall.SockaddrNetlink{
		Family: syscall.AF_NETLINK,
		Pid:    uint32(0),
		Groups: uint32(groups),
	}

	err = syscall.Bind(s, saddr)
	if err != nil {
		return nil, fmt.Errorf("bind: %s", err)
	}

	return &NetlinkListener{fd: s, sa: saddr}, nil
}

func main() {
	l, _ := ListenNetlink()
	for {
		_, err := l.ReadMsgs()
		if err != nil {
			fmt.Println("=========   ", err)
		}
	}
}
