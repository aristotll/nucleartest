#include <sys/epoll.h>
#include <unistd.h>
#include <stdlib.h>
#include <string.h>
#include <stdio.h>
#include <errno.h>

#define EVENTS_SIZE 128
#define BUF_SIZE 1024

int main() {
	int pipefd[2];
	int epfd = epoll_create(1024);
	struct epoll_event event1, event2, events[EVENTS_SIZE];
	char buf[BUF_SIZE];

	event1.events = EPOLLIN;
 	event1.data.fd = STDIN_FILENO;
	int ret = epoll_ctl(epfd, EPOLL_CTL_ADD, STDIN_FILENO, &event1);
	if (ret != 0 ) {
		printf("epoll_ctl error: %s\n", strerror(errno));
		return 0;
	}

	if (pipe(pipefd) == -1) {
		perror("pipe");
		exit(EXIT_FAILURE);
	}

	// 注意要先调用 pipe() 创建管道，再添加到 epoll，否则报错 Bad file descriptor
	event2.events = EPOLLIN;
	event2.data.fd = pipefd[0];
	ret = epoll_ctl(epfd, EPOLL_CTL_ADD, pipefd[0], &event2); // 将管道的读端添加到 epoll 监听队列
	if (ret != 0) {
		printf("epoll_ctl error: %s\n", strerror(errno));
		return 0;
	}


	pid_t cpid = fork();

	switch (cpid) {
	case -1:
		perror("fork");
		exit(EXIT_FAILURE);
	case 0:	// child
		// 发现一个新坑：如果使用 printf 进行调试，必须添加 \n 换行符，否则终端不会输出
		// 搞得我还以为子进程没有执行
		// printf("child in\n");
		close(pipefd[0]);	// 关闭管道的读端，因为子进程只负责向管道写入数据
		sleep(3);
		char *s = "break";
		write(pipefd[1], s, strlen(s));
		printf("child send sign to pipe\n");
		close(pipefd[1]);
		exit(0);	// 必须要写这句，不然会出现程序已经退出（终端输出完毕，已经显示新行），
				// 但是进程依然在执行的迷惑情况（终端输入全部会被当成 epoll 事件）
	default: // parent
		close(pipefd[1]);	// 关闭管道的写端，因为父进程只负责从管道中读取
		while (1) {
			printf("epoll is wait\n");
			int n = epoll_wait(epfd, events, EVENTS_SIZE, -1); // 无限阻塞
			for (int i = 0; i < n; i++) {
				//printf("event fd: %d\n", events[i].data.fd);
				if (events[i].data.fd == STDIN_FILENO) {
					memset(buf, '\0', BUF_SIZE);
					read(events[i].data.fd, buf, BUF_SIZE);
					printf("read from stdin: %s\n", buf);
				} else if (events[i].data.fd == pipefd[0]) {
					printf("pipe break epoll_wait!\n");
					memset(buf, '\0', BUF_SIZE);
					read(events[i].data.fd, buf, BUF_SIZE);	// 消费掉事件
					close(pipefd[0]);
					goto STOP_WAIT;
				}
			}
		}
STOP_WAIT:
		printf("do something\n");
		exit(0);
	}
}
