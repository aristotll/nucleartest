#include <stdio.h>

int main(int argc, char const *argv[]) {
    extern void publicFunc();
    publicFunc();

    extern void privateFunc();
    privateFunc();
    return 0;
}
