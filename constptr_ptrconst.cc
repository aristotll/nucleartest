#include <iostream>

int main() {
    int i = 123, j = 456;
    // 常量指针,修饰的是 int*
    const int* ptr = &i;
    ptr = &j;  // 可以修改指向
    //*ptr = 111; // 但是不可以修改指向的值

    // 指针常量,修饰的是 ptr1
    int* const ptr1 = &i;
    *ptr1 = 111;    // 可以修改指向的值
    //ptr1 = &j;  // 不可以修改指向
}
