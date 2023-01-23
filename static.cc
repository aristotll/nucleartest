#include <iostream>

class Todo {
private:
    static int static_x;
    int y_;
public:
    int getStaticX();
    int getY();
};

int Todo::static_x = 10086;

int Todo::getStaticX() {
    return this->static_x;
}

int Todo::getY() {
    return this->y_;
}

class Dota {
public:
    Dota() = default;
    Dota(int x);
    int getX();
private:
    int x_;
};

Dota::Dota(int x) {
    this->x_ = x;
}

int Dota::getX() {
    return this->x_;
}

int main() {
    Todo* todo = new Todo();
    std::cout << todo->getStaticX() << std::endl;
    delete todo;

    std::cout << todo->getY() << std::endl;
    
    Dota dota(todo->getStaticX());
    std::cout << dota.getX() << std::endl;
}
