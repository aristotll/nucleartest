#include <signal.h>
#include <stdio.h>
#include <unistd.h>
#include <stdlib.h>

void handler(int sig) { printf("SIGCHLD coming! \n"); }

int main(int argc, char const *argv[]) {
    pid_t pid;
    sigset_t mask_one, prev_one;
    signal(SIGCHLD, handler);
    sigemptyset(&mask_one);
    sigaddset(&mask_one, SIGCHLD);
    sigprocmask(SIG_BLOCK, &mask_one, &prev_one);

    if ((pid = fork()) == 0) {  // child
        //sigprocmask(SIG_SETMASK, &prev_one, NULL); // unlock SIGCHLD
        exit(0);
    }

    sleep(5);
    sigprocmask(SIG_SETMASK, &prev_one, NULL); // restore

    sleep(15);
    return 0;
}
