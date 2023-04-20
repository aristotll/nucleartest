#include <set>
#include <string>
#include <iostream>

class Person {
public:
    std::string name;
    int age;
    Person(const std::string& n, int a) : name(n), age(a) {}

    // 重载指针所指向对象的 operator< 运算符
    bool operator<(const Person& other) const {
        if (name != other.name) {
            return name < other.name;
        }
        return age < other.age;
    }
};

int main() {
    // 创建 set 容器，并插入元素
    std::set<Person*> personSet;
    personSet.insert(new Person("Alice", 25));
    personSet.insert(new Person("Bob", 30));
    personSet.insert(new Person("Alice", 25)); // 不会插入，因为已经存在相同的元素

    // 输出 set 容器中的元素数量
    std::cout << "set size: " << personSet.size() << std::endl;

    // 释放指针所指向的内存空间
    for (Person* p : personSet) {
        delete p;
    }
    return 0;
}
