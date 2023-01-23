#include <iostream>

class A {
private:
    int a = 0;
    int b = 0;
public:
    A(int a_) : a(a_) {};
    A(int a_, int b_) : a(a_), b(b_) {} ;
    friend std::ostream& operator<<(std::ostream& out, const A& a) {
        return out << "a: " << a.a << " b: " << a.b;
    }
};

int main() {
    A a = 10086;
    std::cout << a << std::endl;
}
