#include <iostream>
#include <stdlib.h>
#include <unistd.h>

using namespace std;

int main(int argc, char const *argv[]) {
    string str = "hello, world \n";
    write(1, str.data(), str.size());
    return 0;
}
