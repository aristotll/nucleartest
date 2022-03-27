package main

import (
	"fmt"
	"strings"
)

func main() {
	s := `40 35 0:34 / /sys/fs/cgroup/cpu,cpuacct rw,nosuid,nodev,noexec,relatime shared:649 - cgroup cgroup rw,cpu,cpuacct`
	ss := strings.Split(s, " ")

	for index, val := range ss {
		fmt.Printf("[%v] -> %v \n", index+1, val)
	} 
	//fmt.Println(ss)

	s = "localhost:8080"
	ss = strings.Split(s, ":")
	fmt.Println(ss)
}
