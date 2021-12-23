#include <signal.h>
#include <stdio.h>
#include <unistd.h>

// 验证：同类型信号下，最多只能有一个待处理信号，之后的会被丢弃
// 流程：先发送第一个信号，之后立马发送第二个信号，再发送第三个信号，
// 看 printf 会输出几次
//
// 结论：3 秒后输出第一个 printf，之后经过 3 秒输出第二个 printf，之后程序结束。
// 整个过程 printf 只会输出两次，这代表最多只能有个一个待处理信号
// 
void handler(int sign) {
    sleep(3);  
    printf("SIGINT is coming! \n");
}

int main(int argc, char const *argv[]) {
    signal(SIGINT, handler);
    sleep(100); 
    return 0;
}
