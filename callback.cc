#include <iostream>

using namespace std;

int sum(int x, int y) {
	return x + y;
}	

int Do(int x, int y, int (*fn)(int, int)) {
	cout << fn(x, y) << endl;
}

int main() {
	Do(1, 2, sum);
}
