#include <arpa/inet.h>
#include <stdio.h>

#include <iostream>

int main() {
	// error
    char dst[sizeof(struct in_addr)];
    const char* src = "128.2.194.242";
    int ret = inet_pton(AF_INET, src, &dst);
    if (ret < 0) {
        std::cout << "error" << std::endl;
    }
    printf("%s\n", dst);
    std::cout << dst << std::endl;

	// true
    struct sockaddr_in sa;
    // store this IP address in sa:
    inet_pton(AF_INET, "128.2.194.242", &(sa.sin_addr));
	std::cout << sa.sin_addr.s_addr << std::endl;
}
