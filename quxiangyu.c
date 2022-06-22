#include <stdio.h>

int main() {
	int x = 10;
	// --> 趋向于（愚人节玩笑）
	while (x --> 0) {
		printf("%d ", x);
	}
	return 0;
}
