#include <iostream>

template <typename T>
T max(T x, T y) {
    return x > y ? x : y;
}

class Class {
   public:
    Class(int x) { this->x = x; };
    int getX() { return this->x; }
	bool operator>(Class &c) {
		return this->x > c.x;
	}

   private:
    int x;
};

int main() {
    std::cout << max(1, 2) << std::endl;
    Class c1(1), c2(2);
    std::cout << max(c1, c2) << std::endl;
}
