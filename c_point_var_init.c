#include <stdio.h>

int main(int argc, char const *argv[]) {
    int *i;
    printf("%p\n", i);
    *i= 5;
    return 0;
}
