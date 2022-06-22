#include <cstdlib>

class Apple {
public:
	Apple() {};
	int getCount() {
		return m_iData;
	};
private:
	int m_iData;
};

int main() {
	const Apple apple;
	apple.getCount();
	return 1;
}
