#include <iostream>

int main() {
    int ival = 1024;
    int &refVal = ival;
    int &refVal1 = refVal;
    refVal1 = 1000;
    std::cout << ival << std::endl;
    int i = refVal;
    i = 1024;
    std::cout << ival << refVal << std::endl;
}
