#include <iostream>
#include <dlfcn.h>
#include <functional>

int main() {
    void *handle = dlopen("./libsum.so", RTLD_NOW);
    if (!handle) {
        std::cerr << dlerror() << std::endl;
        return 1;    
    }
    //typedef int(*func)(int, int);
    using func = int (*)(int, int);
    //int (*sum)(int, int);
    //func sum;
    func sum = (func)dlsym(handle, "sum");
    if (!sum) {
        std::cerr << dlerror() << '\n';
        dlclose(handle);
        return 1;
    }
    int a = 1;
    int b = 2;
    std::cout << a << " + " << b << " = " << sum(a, b)
    << std::endl;

    dlclose(handle);
    return 0;
}
