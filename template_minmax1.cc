#include <iostream>

//using namespace std;

class A {};
class B {};

template <typename T>
T max(const T& x, const T& y) {
    return x > y ? x : y;
}

template <class T>
T max1(const T& x, const T& y) {
    return x > y ? x : y;
}

int main() {
    int a = max<int>(1, 2);
    bool b = max<bool>(true, false);
    A a1();
    B b1();
    max1<class>(a1, b1);
    std::cout << a << " " << b << std::endl;
}
