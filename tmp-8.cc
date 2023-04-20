#include <iostream>

void checkint(int i);

int main() {
    checkint(0);
    checkint(1);
    checkint(123);
}

void checkint(int i) {
    if (i)
        std::cout << 1 << std::endl;
    else 
        std::cout << 0 << std::endl;
}
