#include <stdio.h>

void mulstore(long, long, long*);

long mult2(long a, long b) {
    long s = a * b;
    return s;
}

int main() {
    long d;
    long x = 2, y = 3;
    mulstore(2, 3, &d);
    printf("%ld * %ld = %ld \n", x, y, d);
    return 0;
}
