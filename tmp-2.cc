#include <iostream>

int main() {
    int i = 111;
    int &ii = i;
    ii = 1.1;
    std::cout << ii << std::endl;
}
