#include <iostream>
#include <unistd.h>
#include <fcntl.h>
#include <errno.h>
#include <strings.h>

using namespace std;

int main() {
	int fd = open("1.txt", O_RDWR|O_CREAT, 0777);
	if (fd == -1) {
		fprintf(stderr, "%s: %s \n", "open file error", strerror(errno));
		return 0;
	}
	printf("open file fd: %d \n", fd);

	pid_t pid = fork();
	if (pid == -1) {
		fprintf(stderr, "%s: %s \n", "fork error", strerror(errno));
		return 0;
	}

	if (pid == 0) {	// child
		string s = "child write: [1, 2, 3] \n";
		int n = write(fd, s.c_str(), s.size());
		if (n == -1) {
			fprintf(stderr, "%s: %s \n", "write file error", strerror(errno));
			return 0;
		}
	} else { //parent
		string s = "parent write: [a, b, c] \n";
		int n = write(fd, s.c_str(), s.size());
		if (n == -1) {
			fprintf(stderr, "%s: %s \n", "write file error", strerror(errno));
			return 0;
		}
	}
}

