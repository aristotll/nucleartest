#include <stdio.h>

#define N 5

void func(int a[]) {
    for (int i = 0; i < N; i++) {
        printf("[%d]: %d \n", i, a[i]);
    }
}

void func1(int a[N]) {
    for (int i = 0; i < N; i++) {
        printf("[%d]: %d \n", i, a[i]);
    }
}

int main(int argc, char const *argv[]) {
    int a[N] = {1, 2, 3, 4, 5};
    func(a);
    func1(a);
    return 0;
}
