#include <errno.h>
#include <stdio.h>
#include <stdlib.h>
#include <sys/wait.h>
#include <unistd.h>

#define N 2

int main(int argc, char const *argv[]) {
    int pstat;
    pid_t pid;

    for (int i = 0; i < N; i++) {
        if ((pid = fork()) == 0) {
            exit(100 + i);
        }
    }

    while ((pid = waitpid(-1, &pstat, 0)) > 0) {
        if (WIFEXITED(pstat)) {
            printf("子进程 %d 正常退出，退出状态：%d \n", pid,
                   WEXITSTATUS(pstat));
        } else {
            printf("子进程 %d 异常终止 \n", pid);
        }
    }

    if (errno != ECHILD) {
        fprintf(stderr, "waitpid error \n");
    }

    return 0;
}
