#include <stdio.h>
#include <stdlib.h>

void printArr(int*, int);

int main(int argc, char const *argv[]) {
    /* code */
    int a[] = {1, 2, 3};
    int *aa = malloc(sizeof(int) * 3);
    aa[0] = 4;
    aa[1] = 5;
    aa[2] = 6;
    
    printArr(a, 3);
    printArr(aa, 3);
    
    return 0;
}

void printArr(int *arr, int len) {
    for (int i = 0; i < len; i++) {
        /* code */
        printf("%d \n", arr[i]);
    }
}
