#include <stdio.h>

#define MAX(x, y) x > y ? x : y

int main() {
	char *a = "123";
	char *b = "456";
	MAX(1, 2);
	MAX(a, b);
}
