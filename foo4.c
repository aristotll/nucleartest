#include <stdio.h>

void f(void);

int x;

int main(int argc, char const *argv[]) {
    x = 15213;
    f();
    printf("x = %d\n", x);
    return 0;
}
