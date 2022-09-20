#include <iostream>
#include <functional>

void func(std::function<int(int, int)> fn, int x, int y) {
    std::cout << fn(x, y) << std::endl;
}

void func1(int(*fn)(int, int), int x, int y) {
    std::cout << fn(x, y) << std::endl;
}

int main() {
    func([](int x, int y) {
        return x + y;
    }, 1, 2);

    func([](int x, int y) {
        return x / y;
    }, 10, 5);

    func([](int x, int y) {
        return x * y;
    }, 2, 30);
}
