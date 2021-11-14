#include <unistd.h>
#include <iostream>
#include <csignal>

using namespace std;

int sleep_ret;

void handler(int sig) {
    cout << "sleep: " << sleep_ret << endl;
}

int main(int argc, char const *argv[]) {
    signal(SIGINT, handler);
    cout << "pid: " << getpid() << endl;
    sleep_ret = sleep(60);
    return 0;
}
