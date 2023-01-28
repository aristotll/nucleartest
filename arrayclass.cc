#include <iostream>

class A {
 public:
     A() {
         //初始化arr
         //arr[0] = 0;
         //arr[1] = 0;
         //arr[2] = 0;
         //arr[3] = 0;
     }
     void printArr() {
        std::cout << "arr[0] -> " << this->arr[0] << std::endl;
        std::cout << "arr[1] -> " << this->arr[1] << std::endl;
        std::cout << "arr[2] -> " << this->arr[2] << std::endl;
        std::cout << "arr[3] -> " << this->arr[3] << std::endl;
    }

 public:
     int arr[4];
 };

int main() {
    A a;
    a.printArr();    
}
