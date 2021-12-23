#include <fcntl.h>
#include <stdio.h>
#include <sys/mman.h>
#include <sys/stat.h>
#include <unistd.h>

int main(int argc, char const *argv[]) {
    char *file_name = argv[1];
    int fd = open(file_name, O_CREAT | O_RDWR, 0777);
    if (fd < 0) {
        fprintf(stderr, "open file error \n");
        return 0;
    }

    write(fd, "123456", 6);
    struct stat sb;
    fstat(fd, &sb);

    char *start =
        mmap(NULL, sb.st_size, PROT_READ | PROT_WRITE, MAP_SHARED, fd, 0);
    if (start == MAP_FAILED) {
        fprintf(stderr, "open file error");
        return 0;
    }

    fprintf(stdout, start);

    return 0;
}
