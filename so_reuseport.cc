#include <arpa/inet.h>
#include <assert.h>
#include <stdio.h>
#include <strings.h>
#include <sys/socket.h>
#include <sys/types.h>
#include <unistd.h>

// g++ -std=c++11 -o so_reuseport so_reuseport.cc && ./so_reuseport

static int tcp_listen(char *ip, int port) {
    int lfd, opt, err;
    sockaddr_in addr;

    lfd = socket(AF_INET, SOCK_STREAM, IPPROTO_TCP);
    assert(lfd != -1);

    opt = 1;
    err = setsockopt(lfd, SOL_SOCKET, SO_REUSEPORT, &opt, sizeof(opt));
    assert(!err);

    bzero(&addr, sizeof(addr));
    addr.sin_family = AF_INET;
    addr.sin_addr.s_addr = inet_addr(ip);
    addr.sin_port = htons(port);

    err = bind(lfd, (sockaddr *)&addr, sizeof(addr));
    assert(!err);

    err = listen(lfd, 8);
    assert(!err);

    return lfd;
}

int main(int argc, char **argv) {
    int lfd, sfd;
    lfd = tcp_listen("127.0.0.1", 8080);
    while (true) {
        sfd = accept(lfd, nullptr, nullptr);
        close(sfd);
        printf("接收到 tcp 连接：%d \n", sfd);
    }

    return 0;
}