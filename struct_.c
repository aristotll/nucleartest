#include <stdio.h>

typedef struct {
	int x;
	int y;
} Struct;

int add(Struct *s) {
	return s->x + s->y;
}

int main() {
	Struct s = {
		.x = 1,
		.y = 2,
	};
	printf("%d\n", add(&s));

	Struct s1 = {
		x: 1,
		y: 2,
	};
	printf("%d\n", add(&s1));
}
