#include <csignal>
#include <iostream>
#include <unistd.h>

using namespace std;

void intHandler(int sign) {
    cout << "SIGINT 处理中..." << endl;
    sleep(10);
    cout << "SIGINT 处理完毕" << endl;
}

void termHandler(int sign) {
    cout << "SIGTERM 处理中..." << endl;
    sleep(5);
    cout << "SIGTERM 处理完毕" << endl;
    sleep(5);
}

// 信号处理程序被其他信号处理程序中断
int main(int argc, char const *argv[]) {
    cout << "pid: " << getpid() << endl;
    signal(SIGINT, intHandler);
    signal(SIGTERM, termHandler);
    sleep(60);
    return 0;
}
