#include <iostream>

using namespace std;

class A {
public:
	string name;
	void setName(string name) const {
		this->name = name;
	}
};

void f(const A& a) {
	a.setName("aaa");
}

int main() {
	
}
