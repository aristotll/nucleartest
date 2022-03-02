#include <stdio.h>

#define N 5

int sumarray3d(int a[N][N][N]) {
	int i, j, k, sum;
	for (i = 0; i < N; i++) {
		for (j = 0; j < N; j++) {
			for (k = 0; k < N; k++) {
				printf("%d ", a[k][i][j]);
				sum += a[k][i][j];
			}
		}
	}
	return sum;
}

int main() {
	int a[N][N][N] = {
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
		{11, 12, 13, 14, 15},
		{16, 17, 18, 19, 20, 21},
		{22, 23, 24, 25, 26, 27},
	};
	printf("%d \n", sumarray3d(a));
}
