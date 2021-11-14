#include <unistd.h>

#include <csignal>
#include <iostream>

using namespace std;

int main(int argc, char const *argv[]) {
    pid_t pid;
    cout << "parent pid: " << getpid() << endl;
    if ((pid = fork()) == 0) {
        cout << "child pid: " << getpid() << endl;
        pause();
        cout << "控制永远不应该到达这里" << endl;
        exit(0);
    }

    kill(0, SIGKILL);
    return 0;
}
