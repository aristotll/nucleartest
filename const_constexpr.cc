#include <iostream>

const int add(const int x, const int y) {
    return x + y;
}

constexpr int mul(const int x, const int y) {
    return x * y;
}

int main() {
    int x, y;
    std::cin >> x;
    std::cin >> y;
    std::cout << mul(x, y);   
}
