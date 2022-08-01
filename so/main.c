#include <stdio.h>
#include <unistd.h>
#include "print.h"

int main() {
	printf("waite 5 seconds\n");
	sleep(5);
	printf("%s\n", __FUNCTION__);
	printExtern();

	return 0;
}
