#include <stdio.h>
#include <stdlib.h>

int main(int argc, char const *argv[]) {
    int *num = (int*) malloc(sizeof(int) * 5);
    for (int i = 0; i < 5; i++) {
        printf("%d \n", num[i]);
    }

    int *num1 = (int*) calloc(5, sizeof(int));
    for (int i = 0; i < 5; i++) {
        printf("%d \n", num1[i]);
    }

    return 0;
}
