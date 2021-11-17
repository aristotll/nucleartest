#include <stdlib.h>
#include <stdio.h>
#include <errno.h>
#include <strings.h>

int main() {
	double loadavg[3];
	int num = getloadavg(loadavg, 3);
	if (num == -1) {
		printf("getloadavg error: %s \n", strerror(errno));
		return 0;	
	}
	for ( int i = 0; i < 3; i++) {
		double use = loadavg[i] * 100 / 8;
		printf("%lf%% \n", use);
	}
}
