#include <stdio.h>

int leaf(int y) {
	return y + 2;
}

int top(int x) {
	int ret = leaf(x-5);
	//printf("%d\n", ret);
	return ret + ret;
}

int main() {
	int ret = top(100);
	printf("%d\n", ret);
}
