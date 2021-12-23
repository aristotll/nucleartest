#include <unistd.h>

#include <cerrno>
#include <csignal>
#include <cstring>
#include <iostream>

using namespace std;

void snooze(int totalTime, int usedTime) {
    printf("Slept for %d of %d secs. \n", usedTime, totalTime);
}

void intHandler(int sig) {
    cout << "catch SIGINT !" << endl;
    return;
}

int main(int argc, char const *argv[]) {
    if (argc != 2) {
        fprintf(stderr, "usage: ./snooze N \n");
        return 0;
    }
    signal(SIGINT, intHandler);

    int totalTime = atoi(argv[1]);
    if (totalTime == 0) {
        fprintf(stderr, "atoi error: %s \n", strerror(errno));
        return 0;
    }

    int sleepRet = sleep(totalTime);
    snooze(totalTime, totalTime - sleepRet);
    return 0;
}
