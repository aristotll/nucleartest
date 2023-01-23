#include <iostream>

class Todo {
public:
    Todo(int x, int y, std::string s);
private:
    int x_;
    int y_;
    std::string s_;
};

Todo::Todo(int x, int y, std::string s) : x_(x), y_(y), s_(s) {}

int main() {
    Todo todo(1, 2, "abc");
}
