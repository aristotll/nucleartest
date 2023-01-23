#include <iostream>
#include <functional>

// 这种写法不支持 c++ 的 lambda 传参
int fn(int x, int y, double (*func)(int x, int y)) {
    return func(x, y);
}

int fn1(double x, double y, std::function<double(double, double)> func) {
    return func(x, y);
}

int main() {
    int ret = fn1(2, 3, [](int x, int y){
        return x + y;
    });
    std::cout << ret << std::endl;
}
