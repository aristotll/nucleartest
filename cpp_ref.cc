#include <stdio.h>

int main() {
	int i = 10;
	int& a = i;
	int& b = i;

	a = 100;
	printf("%d\n", b);
}
