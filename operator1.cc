#include <iostream>

class Todo {
private:
    int x_;
public:
    Todo(int x);
    // 成员重载运算符, this 会作为第一个运算参数
    // 比如这里重载了 + 运算符,那么 this 就会默认作为 
    // 左边的运算数, 参数 other 作为右边的运算数
    int operator+(Todo other);
};

Todo::Todo(int x) {
    this->x_ = x;
}

int Todo::operator+(Todo other) {
    return this->x_ + other.x_; 
}

int main() {
    Todo todo1(10);
    Todo todo2(20);
    //int res = operator+(1, 2);
    std::cout << todo1.operator+(todo2) << std::endl;
    std::cout << todo1 + todo2 << std::endl;
}
