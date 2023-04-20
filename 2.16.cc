#include <iostream>

int main() {
    int i = 0, &r1 = i;
    double d = 0, &r2 = d;

    r2 = 3.1415;
    std::cout << r2 << std::endl;
    r2 = r1;
    i = r2;
    r1 = d;

    std::cout << r2 << " " << i << " " << r1 << std::endl;
}
