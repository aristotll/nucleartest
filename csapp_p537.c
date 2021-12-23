#include <errno.h>
#include <signal.h>
#include <stdio.h>
#include <stdlib.h>
#include <strings.h>
#include <sys/wait.h>
#include <unistd.h>

void handler1(int sig) {
    int olderrno = errno;

    if ((waitpid(-1, NULL, 0)) < 0) {
        fprintf(stderr, "waitpid error \n");
        exit(0);
    }
    char *errmsg = "Handler reaped child \n";
    write(STDOUT_FILENO, errmsg, strlen(errmsg));
    sleep(1);
    errno = olderrno;
}

int main(int argc, char const *argv[]) {
    int i, n;
    char buf[4096];

    if (signal(SIGCHLD, handler1) == SIG_ERR) {
        fprintf(stderr, "signal error");
        exit(0);
    }

    for (int i = 0; i < 3; i++) {
        if (fork() == 0) {
            printf("Hello from child %d\n", getpid());
            exit(0);
        }
    }

    if (n = read(STDIN_FILENO, buf, sizeof(buf)) < 0) {
        fprintf(stderr, "read error");
        exit(0);
    }

    printf("Parent processing input \n");
    while (1);
    
    return 0;
}
