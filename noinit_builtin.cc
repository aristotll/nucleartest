#include <iostream>

int global_int;
std::string global_str;

int main() {
    int a, d;
    size_t b;
    double c;
    std::string str;
    static int f;

    std::cout << a << std::endl;
    std::cout << b << std::endl;
    std::cout << c << std::endl;
    std::cout << d << std::endl;
    std::cout << global_int << std::endl;
    std::cout << global_str << std::endl;
    std::cout << str << std::endl;
    std::cout << f << std::endl;
}
