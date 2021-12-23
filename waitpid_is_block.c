#include <errno.h>
#include <stdio.h>
#include <stdlib.h>
#include <sys/wait.h>
#include <unistd.h>

int main(int argc, char const *argv[]) {
    pid_t pid;
    for (int i = 0; i < 5; i++) {
        if ((pid = fork()) == 0) {  // child
            printf("child[%d] is working... \n", getpid());
            sleep(i + 1);
            printf("child[%d] is done! \n", getpid());
            exit(0);
        }
    }

    int stat;
    //waitpid(-1, &stat, 0);
    while (waitpid(-1, &stat, 0) > 0) {

    }

    if (errno != ECHILD) {
        fprintf(stderr, "waitpid error \n");
    }

    printf("parent process is done! \n");

    return 0;
}
