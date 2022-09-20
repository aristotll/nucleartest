#include <iostream>
#include <arpa/inet.h>
#include <sys/socket.h>
#include <string.h>
#include <thread>
#include <strings.h>

#define PORT 8080
#define BACKLOG_SIZE 10

int main() {
    int iSocketFD = socket(AF_INET, SOCK_STREAM, 0); // 建立socket
	if (0 > iSocketFD) {
		printf("创建socket失败！\n");
		return 0;
	}

    struct sockaddr_in stLocalAddr;
    struct sockaddr_in stRemoteAddr;
    bzero(&stLocalAddr, sizeof(stLocalAddr));
    bzero(&stRemoteAddr, sizeof(stRemoteAddr));

    stLocalAddr.sin_family = AF_INET;  /*该属性表示接收本机或其他机器传输*/
	stLocalAddr.sin_port = htons(PORT); /*端口号*/
	stLocalAddr.sin_addr.s_addr=htonl(INADDR_ANY); /*IP，括号内容表示本机IP*/

    //绑定地址结构体和socket
	if (0 > bind(iSocketFD, (struct sockaddr*) &stLocalAddr, sizeof(stLocalAddr))) {
		printf("绑定失败！\n");
		return 0;
    }

    //开启监听 ，第二个参数是最大监听数
	if (0 > listen(iSocketFD, BACKLOG_SIZE)) {
		printf("监听失败！\n");
		return 0;
	}

	printf("iSocketFD: %d\n", iSocketFD);

    socklen_t sin_size = sizeof(struct sockaddr_in);
    //在这里阻塞知道接收到消息，参数分别是socket句柄，接收到的地址信息以及大小 
	int new_fd = accept(iSocketFD, (struct sockaddr*)&stRemoteAddr, &sin_size);
	if( 0 > new_fd) {
		printf("接收失败！\n");
		return 0;
	} else {
		printf("接收成功！\n");
        for (int i = 0; i < 100; i++) {
            std::thread th([new_fd, i]() {
                if (i%2 == 0) {
                    send(new_fd, "aaa\n", sizeof("aaa\n"), 0);
                } else {
                    send(new_fd, "bbb\n", sizeof("bbb\n"), 0);
                }
            });
            th.join();
        }
	}
}

void sendFunc(int n) {
    
}
