#include <iostream>

class Todo {
private:
    int x;
    int y;
public:
    int getX() const;
    int getY();
    void setX(int x);
    void setY(int y);
};

int Todo::getX() const {
   // this->x++;
    return this->x;
}

int Todo::getY() {
    this->y++;
    return this->y;
}

void Todo::setX(int x) {
    this->x = x;
}

void Todo::setY(int y) {
    this->y = y;
}

int main() {
    Todo todo;
    todo.setX(1);
    todo.setY(1);
    std::cout << "x: " << todo.getX() << " y: " << todo.getY() << std::endl;
}
