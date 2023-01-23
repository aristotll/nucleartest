#include <iostream>

constexpr int add(int x, int y) {
    return x + y;
}

void fn() {
    
}

constexpr float exp(float x, int n) {
    return n == 0 ? 1 :
        n % 2 == 0 ? exp(x * x, n / 2) :
        exp(x * x, (n - 1) / 2) * x;
}

int main() {
    const int res = add(1, 2);
    //res = 1;
    constexpr int res1 = add(1, 2);
    constexpr int x = 123;
    constexpr int res2 = add(x, 1);
    std::cout << res << std::endl;
    std::cout << exp(1, 10) << std::endl;
}
