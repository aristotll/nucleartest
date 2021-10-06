#include <iostream>

using namespace std;

class Interface {
public:
    virtual void open() = 0;
    virtual void close() = 0;
};

class Obj : public Interface {
public:
    Obj(/* args */) {}
    ~Obj() {}

    void open() {
        cout << "open" << endl;
    }

    void close() {
        cout << "close" << endl;
    }
};

void func(Interface *i) {
    i->open();
    i->close();
}

int main(int argc, char const *argv[]) {
    Interface *a = new Obj();
    func(a);
    return 0;
}
