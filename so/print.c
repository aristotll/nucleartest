#include <stdio.h>

void printInter() {
	printf("%s\n", __FUNCTION__);
}

void printExtern() {
	printInter();
	printf("%s\n", __FUNCTION__);
}
