#include <unistd.h>

#include <csignal>
#include <iostream>

using namespace std;

void kill_handler(int sign) {
    cout << "接收到 kill 信号" << endl;
    cout << "保存中..." << endl;
}

void int_handler(int sign) {
    cout << "接收到 SIGINT 信号" << endl;
    exit(0);
}

void term_handler(int sign) {
    cout << "接收到 SIGTERM 信号" << endl; 
    cout << "保存中..." << endl;
    exit(0);
}

int main(int argc, char const *argv[]) {
    cout << "pid: " << getpid() << endl;
    signal(SIGKILL, kill_handler);
    signal(SIGINT, int_handler);
    signal(SIGTERM, term_handler);
    sleep(60);
    return 0;
}
