#include <stdio.h>

int *_xx() {
	static int y = 10;
	printf("%d\n", y);
	return &y;	// 可以正常返回指针，因为这是static 变量
}

int *__xx() {
	int y = 10;
	printf("%d\n", y);
	return &y;	// 返回局部变量的指针，会产生警告 
}

int main() {
	int *ret = _xx();
	printf("%d\n", *ret);
	(*ret)++;
	_xx();

	printf("===================\n");
	int *ret1 = __xx();
	printf("%d\n", *ret1);
	(*ret1)++;
	__xx();
}
