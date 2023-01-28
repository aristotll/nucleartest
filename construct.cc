#include <iostream>

class Class {
public:
    Class(int x) : x_(x) {}
    void printX(); 
private:
    int x_;
};

void Class::printX() {
    std::cout << x_ << std::endl;
}

int main() {
    Class c = 1;
    c.printX();
}
