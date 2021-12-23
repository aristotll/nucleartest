#include <signal.h>
#include <stdio.h>
#include <unistd.h>

void handler_parent(int sig) { printf("[parent] SIGINT coming \n"); }

void handler_child(int sig) { printf("[child] SIGINT coming \n"); }

int main(int argc, char const *argv[]) {
    pid_t pid;

    if ((pid = fork()) == 0) {  // child
        signal(SIGINT, handler_child);
        sleep(10);
    }
    signal(SIGINT, handler_parent);

    sleep(20);
    return 0;
}
