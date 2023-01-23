#include <iostream>

using fun = void (*)(int, int);

void fn(int x, int y, fun f) {
    f(x, y);
}

void fn1(int x, int y) {
    std::cout << x + y << std::endl;
}

int main() {
    fn(1, 2, fn1);
}
