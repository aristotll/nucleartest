#include <stdio.h>
#include <unistd.h>
#include <strings.h>

int main(int argc, char const *argv[]) {
    char *str = "hello, world \n";
    write(1, str, strlen(str));
    return 0;
}