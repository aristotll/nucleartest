#include <iostream>
#include <dlfcn.h>

#define DL_LIB_PATH "./libcaculate.so"
//using FUNC = int (*) (int, int);
typedef int (*CAC_FUNC)(int, int);

int main() {
    CAC_FUNC cac_func = NULL;
    void *handle = dlopen(DL_LIB_PATH, RTLD_LAZY);
    if (!handle) {
        std::cerr << dlerror() << std::endl;
        return 1;
    }
    // 清除之前存在的错误
    //dlerror();

    char *error;
    *(void **) (&cac_func) = dlsym(handle, "add");
    //*(void **) (&fn) = dlsym(handle, "add");
    if ((error == dlerror()) != NULL) {
        std::cerr << error << std::endl;
        return 1;
    }
    std::cout << (*cac_func)(1, 2) << std::endl;
    dlclose(handle);
}
