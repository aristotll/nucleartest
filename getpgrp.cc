#include <unistd.h>
#include <iostream>

using namespace std;

int main(int argc, char const *argv[]) {
    cout << "process group id: " << getpgrp() << endl;
    if (fork() == 0) {
        cout << "process group id: " << getpgrp() << endl;
        exit(0);
    }
    return 0;
}
