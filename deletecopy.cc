#include <iostream>

class Todo {
private:
    int id_;
public:
    Todo(int id);
    Todo(const Todo& oth) = delete;
    friend void fn(Todo& todo);
    friend void deepCopy(Todo todo);
};

Todo::Todo(int id) : id_(id) {}

void fn(Todo& todo) {
    std::cout << todo.id_ << std::endl; 
}

void deepCopy(Todo todo) {
    std::cout << todo.id_ << std::endl;
}

int main() {
    Todo todo(1);
    fn(todo);
    deepCopy(todo);
}
