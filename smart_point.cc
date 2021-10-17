#include <iostream>
#include <memory>


void sharedPtr(std::shared_ptr<int> i) { (*i)++; }

int main(int argc, char const *argv[]) {
    auto i = std::make_shared<int>(10);
    std::cout << "before value: " << *i << std::endl;
    sharedPtr(i);
    std::cout << "after value: " << *i << std::endl;

    auto nativePtr = i.get();
    std::cout << *nativePtr << std::endl;

    std::cout << i.use_count() << std::endl;    // 1
    auto ii = i;    // 引用 i
    std::cout << i.use_count() << std::endl;    // 2
    auto iii = i;
    i.reset(); // 引用计数重置为 0
    std::cout << i.use_count() << std::endl;    // 0
    // 引用计数为 0 后会释放该变量的内存，此时再访问会产生错误 segmentation fault
    std::cout << *i << std::endl;   

    int *a = new int(1);

    return 0;
}
