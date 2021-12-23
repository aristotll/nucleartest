#include <iostream>

using namespace std;

int main(int argc, char const *argv[]) {
    int i = 10;
    int a = i;

    printf("i = %d \n", a);

    __asm {
        mov dword ptr [ebp-4], 20h
    }

    int b = i;
    printf("i = %d \n", b);
    return 0;
}
