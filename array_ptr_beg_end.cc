#include <iostream>

#define SIZE 5

void printarray(const int* beg, const int* end) {
    while (beg != end) {
        //std::cout << beg << std::endl;
        std::cout << *beg++ << std::endl;
    }
}

int main() {
    int j[SIZE] = {1, 2, 3, 4, 5};
    int* beg = j;
    int* end = j + SIZE;
    printarray(beg, end);
}
