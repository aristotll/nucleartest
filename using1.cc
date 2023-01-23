#include <iostream>

class MyClass {
public:
static const int myInt = 10;
};

using MyInt = MyClass::myInt;

int main() {
std::cout << MyInt << std::endl; // 输出 10
return 0;
}
