package main

import (
	"os"
	"fmt"
	"syscall"
	"os/exec"
	"log"
	"path"
	"strconv"
	"io/ioutil"
)

const cgroupMemoryHierarchyMount = "/sys/fs/cgroup/memory"

func main() {
	fmt.Println("main")
	//fmt.Println("os.Args[0] = ", os.Args[0])
	//fmt.Printf("os.Args[0] == /proc/self/exe ? %v \n", os.Args[0] == "/proc/self/exe")
	// 第一次不会执行到该分支
	if os.Args[0] == "/proc/self/exe" {
		//fmt.Println("os.Args[0] = ", os.Args[0])
		//fmt.Println("os.Args[0] == /proc/self/exe ?")
		// 容器进程
		fmt.Printf("current pid: %v \n", syscall.Getpid())
		// 启动一个占用内存 200m 的进程
		// 该操作会阻塞进程
		cmd := exec.Command("sh", "-c", `stress --vm-bytes 200m --vm-keep -m 1`)
		cmd.SysProcAttr = &syscall.SysProcAttr{}
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	// 执行 /proc/self/exe 会运行当前进程，也就是再次运行该程序 
	cmd := exec.Command("/proc/self/exe")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | 
			   syscall.CLONE_NEWPID |
			   syscall.CLONE_NEWNS,
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Start() 会启动 cmd，并且不等待其结束
	// 如果调用 Run() 或者 CombinedOutput() (内部也是调用了 Run)
	// 那么整个程序会阻塞，因为执行/proc/self/exe 后永远不会结束 
	if err := cmd.Start(); err != nil {
		log.Fatalln(err)
	}
	
	// 得到 fork 出来的进程映射在外部命名空间的 pid	
	fmt.Printf("cmd.Process.Pid = %v \n", cmd.Process.Pid)

	// 在系统默认创建挂载了 memory subsystem 的 hierarchy 上创建 cgroup
	os.Mkdir(path.Join(cgroupMemoryHierarchyMount, "testmemorylimit"), 0755)

	// 将容器进程加入到这个 cgroup 中
	ioutil.WriteFile(path.Join(cgroupMemoryHierarchyMount, "testmemorylimit", "tasks"), []byte(strconv.Itoa(cmd.Process.Pid)), 0644)	

	// 限制 cgroup 进程使用，这里限制进程最多只能使用 100m 内存，而上面的 stress 设置
	// 为启动一个内存占用为 200m 的进程，通过观察 top 查看最终结果
	ioutil.WriteFile(path.Join(cgroupMemoryHierarchyMount, "testmemorylimit", "memory.limit_in_bytes"), []byte("100m"), 0644)

	cmd.Process.Wait()
}
