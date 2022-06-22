#include <iostream>

using namespace std;

class Person {
   private:
    int age;

   public:
    Person() {}
    Person(int age) { this->age = age; }
    Person(const Person& p) {
        age = p.age;
        cout << "拷贝构造函数" << endl;
    }

    Person& operator=(const Person& p) {
        cout << "赋值函数" << endl;
        return *this;
    }
    int getAge() { return this->age; }
    // Person& operator<<(const Person& p) {
    //     cout << p.age << endl;
    //     return *this;
    // }
};

void f(Person p) { return; }

Person f1() {
    Person p(12);
    return p;
}

int main() {
    Person p;
    Person p1 = p;  // 1 拷贝构造函数
    Person p2;
    p2 = p;  // 2 赋值函数
    f(p2);   // 3 拷贝构造函数

    p2 = f1();  // 4 赋值函数

    Person p3 = f1();  // 5 这里有点奇怪，没有输出
    // 需要说明的是，有些编译器出于程序执行效率的考虑，编译的时候进行了优化，
    // 函数返回值对象就不用复制构造函数初始化了，这并不符合 C++ 的标准。
    // 上面的程序，用 Visual Studio 2010 编译后的输出结果如上所述，但是在 Dev
    // C++ 4.9 中不会调用复制构造函数。 把第 14 行的 a
    // 变成全局变量，才会调用复制构造函数。对这一点，读者不必深究。

    cout << f1().getAge() << endl;

    return 0;
}
