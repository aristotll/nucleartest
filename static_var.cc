#include <iostream>

using namespace std;

static int a = 10;

void fn1() {
    int a = 1;
    cout << "fn1 a: " << a << endl;
    a++;
}

void fn2() {
    static int a = 1;
    cout << "fn2 a: " << a << endl;
    a++;
}

int main(int argc, char const *argv[]) {
    /* code */
    for (size_t i = 0; i < 5; i++)
    {
        fn1();
        fn2();
    }
    cout << a++ << endl;
    cout << a << endl;

    return 0;
}
