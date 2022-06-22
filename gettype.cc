#include <iostream>
#include <typeinfo>

int main() {
	std::cout << typeid(1).name() << std::endl;
	std::cout << typeid("1").name() << std::endl;
	std::cout << typeid(1.234).name() << std::endl;
}
