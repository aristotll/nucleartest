#include <stdio.h>
#include <unistd.h>
#include <stdlib.h>
#include <string.h>
#include <sys/wait.h>

int main() {
    int pipefd[2];
    pid_t cpid;

    if (pipe(pipefd) == -1) {
        perror("pipe");
        exit(EXIT_FAILURE);
    }

    cpid = fork();
    switch (cpid) {
    case -1:
        perror("fork");
        exit(EXIT_FAILURE);
    case 0: // child
        close(pipefd[0]);
        char *s = "123\n";
        write(pipefd[1], s, strlen(s));
        close(pipefd[1]);
        _exit(EXIT_SUCCESS);
    default:
        close(pipefd[1]);
        char buf[1024];
        read(pipefd[0], &buf, 1024);
        printf("%s\n", buf);
        close(pipefd[0]);
        wait(NULL);
        _exit(EXIT_SUCCESS);
    }
}
