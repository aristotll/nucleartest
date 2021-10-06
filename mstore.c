#include <stdio.h>

long mult2(long, long);

void mulstore(long x, long y, long *d) {
    long t = mult2(x, y);
    *d = t;
}