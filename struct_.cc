#include <iostream>

typedef struct struct_ {
    int x;
    int y;

    struct_() {}
    struct_(int x, int y) : x(x), y(y) {}
    int add() {return this->x + this->y;}
} Struct;

//int Struct::add() { return this->x + this->y; }

int main() {
    Struct s();

    Struct s1(1, 2);
	std::cout << s1.add() << std::endl;
}