#include <iostream>

using namespace std;

void useNullPtr() {
    char *a;
    *a = '10';
    cout << *a << endl;

    // Output:
    // [1]    4365 segmentation fault  ./segmentation_fault
}

void changeConstVal() {
    const int a = 1;
    a = 50; 
}

int main(int argc, char const *argv[]) {
    /* code */
    return 0;
}
