#include <iostream>

int main() {
    int* i = new int;
    std::cout << *i << std::endl;
    *i = 123;
    std::cout << *i << std::endl;

    int* ii = new int();    
    std::cout << *ii << std::endl;

    int* iii = new int(100);
    std::cout << *iii << std::endl;
}
