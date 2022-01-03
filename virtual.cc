#include <iostream>

using namespace std;

class Interface {
public:
	virtual void A() {
		cout << "parent A()" << endl;
	};
	virtual void B() = 0;
};

class Impl : public Interface {
public:
	virtual void A() {
		cout << "impl A()" << endl;
	}

	void B() {
		cout << "impl B()" << endl;
	}
};

int main() {
	Impl impl;
	impl.A();
	impl.B();
}
