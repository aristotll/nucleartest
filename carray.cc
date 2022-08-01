#include <stdio.h>

void fn(int* a);

int main() {
	int a[] = {1, 2, 3, 4, 5};
	fn(a);
}

void fn(int* a) {
	printf("%d\n", *(a+2));
	*(a+2) = 333;
	printf("%d\n", a[2]);
}
