#include <iostream>

int main() {
    int i = 123;
    int &j = i;
    int *k = j;
    *k = 111;
    std::cout << *k << std::endl;
}
