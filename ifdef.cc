#include <iostream>

#define JOIN_STR(x, y) (std::string)(x##y)
#define CONS(a,b) (int)(a##e##b)
#define ARR_SIZE(a) (sizeof((a))/sizeof((a[0])))

int main() {
    #ifdef JOIN_STR_ERR
        std::cout << JOIN_STR(123, 456) << std::endl;
    #endif

    #ifdef CONS
        std::cout << CONS(2, 3) << std::endl;
    #endif

    int arr[] = {1, 2, 3};
    #ifdef ARR_SIZE
        std::cout << ARR_SIZE(arr) << std::endl;        
    #endif
}
