#include <signal.h>
#include <stdio.h>
#include <unistd.h>

int count;

void handler(int sign) {
    printf("SIGINT[%d] is coming! \n", count++);
}

int main(int argc, char const *argv[]) {
    sigset_t mask, prev_mask;
    sigemptyset(&mask);
    sigaddset(&mask, SIGINT);
    sigprocmask(SIG_BLOCK, &mask, &prev_mask);
    signal(SIGINT, handler);

    sleep(5); // 5 秒后取消阻塞
    // 疑问：如果阻塞时发送多个信号，在取消阻塞后，只会处理一个信号
    sigprocmask(SIG_SETMASK, &prev_mask, NULL);

    sleep(60);
    return 0;
}
