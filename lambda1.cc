#include <iostream>

int main() {
    int i = 10;
    auto func = [](int i) -> int {
        return i * i;
    };
    std::cout << func(i) << std::endl;

    auto func1 = [i] {
        return i * i;
    };
    std::cout << func1() << std::endl;

    auto func2 = [](int i) {
        return i * i;
    };
    std::cout << func2(5) << std::endl;
}
