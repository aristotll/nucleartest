#include <iostream>

int main() {
    int* i;
    {
        int j = 10;
        i = &j;
    }
    std::cout << *i << std::endl;
}
