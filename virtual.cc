#include <iostream>

using namespace std;

class Interface {
public:
	virtual void A() {
		cout << "parent A()" << endl;
	};
	virtual void B() = 0;
    void C() {  // 不加 virtual 关键字，子类也能调用
        cout << "parent C()" << endl;
    };
};

class Impl : public Interface {
public:
	virtual void A() {
		cout << "impl A()" << endl;
	}

	void B() {
		cout << "impl B()" << endl;
	}

    void C() {
        cout << "impl C()" << endl;
    }
};


class Interface1 {
public:
    void a() {  // a 没有加 virtual，但是子类依然可以调用，那加 virtual 的作用在哪里呢？
        cout << "parent a()" << endl;
    }
    virtual void b() {
        cout << "parent b()" << endl;
    }
};

class Impl1 : public Interface1 {
public:
    void b() {
        cout << "child b()" << endl;
    }
};

void itf1Do(Interface1 *i) {
    i->b();
}

int main() {
	Impl impl;
	impl.A();
	impl.B();
    impl.C();

    // 1. 基类函数没加 virtual，子类有相同函数，实现的是覆盖。
    // 用基类指针调用时，调用到的是基类的函数；
    // 用子类指针调用时，调用到的是子类的函数。
    Impl1 *impl1 = new Impl1();
    impl1->a();  // 虽然 a() 没有 virtual，但是子类依然可以调用
    

    // 2. 基类函数加了 virtual 时，实现的是重写。
    // 用基类指针或子类指针调用时，调用到的都是子类的函数。
    // 扯淡？这两输出明明不一样
    impl1->b(); // child b
    Interface1 *itf1 = new Interface1();
    itf1->b();  // parent b
    itf1Do(impl1);
    itf1Do(itf1);

    // 3. 函数加上override，强制要求基本相同函数需要是虚函数，否则会编译报错。
    // 子类的virtual可加可不加，建议加override不加virtual。

}


