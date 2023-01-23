#include <iostream>

class Todo {
public:
    int x_;
    double y_;
    static int z_;
};

int Todo::z_ = 100;

int main() {
    Todo todo;
    std::cout << todo.x_ << ", "<< todo.y_ << ", " << todo.z_ << ", " << std::endl;
}
