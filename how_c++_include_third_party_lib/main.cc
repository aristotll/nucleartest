#include <iostream>
#include "thirdparty/sum.cc"

int main() {
    int a = 1;
    int b = 2;
    std::cout << a << " + " << b << " = " << sum(a, b) 
    << std::endl;
}
