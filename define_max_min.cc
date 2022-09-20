#include <iostream>

#define max(x, y) (x > y ? x : y)

int main() {
	int m = max(1, 2);
	std::cout << max(1, 2) << std::endl;
    std::cout << max("a", "b") << std::endl;
}
