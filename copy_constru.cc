#include<iostream>
using namespace std;
class A {
public:
    int v;
    A(int n) { v = n; };
    A(const A & a) {
        v = a.v;
        cout << "Copy constructor called" << endl;
    }
};

A Func() {
    A a(4);
    return a;
}

int main() {
    cout << Func().v << endl;
    return 0;
}
