#include <iostream>

struct User {
    short age;
    std::string name;
};

int main() {
    auto s = &User{
        .age = 18,
        .name = "abc",
    };
}
