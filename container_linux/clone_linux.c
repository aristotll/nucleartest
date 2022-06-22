#define _GNU_SOURCE
#include <sys/types.h>
#include <sys/wait.h>
#include <stdio.h>
#include <sched.h>
#include <signal.h>
#include <unistd.h>
#define STACK_SIZE (1024 * 1024)

// 函数原型：
// int clone(int (*fn)(void *), void *stack, int flags, void *arg, ...
//     /* pid_t *parent_tid, void *tls, pid_t *child_tid */ );
//
// fn: 当使用clone()创建子进程时，子进程会执行入参的函数fn，clone()的入参arg作为函数fn的参数。
//

static char container_stack[STACK_SIZE];

int container_main(void *args) {
	printf("在容器进行中! \n");
	// 设置 hostname，以验证 CLONE_NEWUTS 是否起作用
	// 第二个参数是第一个参数的长度
	sethostname("container", 9);
	execv("/bin/bash", NULL); // 执行 /bin/bash
} 

int main() {
	printf("程序开始 \n");
	int container_pid = clone(container_main, container_stack+STACK_SIZE, SIGCHLD|CLONE_NEWUTS, NULL);
	// 等待容器进程结束
	waitpid(container_pid, NULL, 0);
	return 0;
}
