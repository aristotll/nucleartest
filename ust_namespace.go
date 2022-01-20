package main

import (
	"os/exec"
	"syscall"
)

func main() {
	cmd := exec.Command("sh")
	cmd.SysProcAttr = &syscall.SysProcAttr{}
	cmd.CombinedOutput()
	cmd.Start()
	cmd.Run()
	syscall.RTNLGRP
}
