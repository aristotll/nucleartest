#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>

int main() {
    pid_t pid = fork();
    int i = 0;
    switch (pid) {
    case -1:
        printf("fork error\n");
        return 0;
    case 0:
        printf("child\n");
    default:
        while(1) {
            printf("parent\n");
            sleep(1);
            //i++;
        }
    }
    return 0;
}
