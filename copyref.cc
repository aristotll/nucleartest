#include <iostream>

int main() {
    int i = 1000;
    int &j = i;
    int &k = j;
    k = 99;
    std::cout << i << std::endl;
    int a = k;
    a = 222;
    std::cout << a << " " << i << std::endl;
}
