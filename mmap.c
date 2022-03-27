#include <fcntl.h>
#include <stdio.h>
#include <sys/mman.h>
#include <sys/stat.h>
#include <unistd.h>

int main(int argc, char const *argv[]) {
    char *file_name = argv[1];
    // 打开文件
    int fd = open(file_name, O_CREAT | O_RDWR, 0777);
    if (fd < 0) {
        fprintf(stderr, "open file error \n");
        return 0;
    }
    // 写入内容到文件
    write(fd, "123456", 6);
    struct stat sb;
    // 获取文件的属性
    fstat(fd, &sb);

    // void *mmap(void *start, size_t length, int prot, int flags,int fd, off_t offset)
    // 
    // start：用户进程中要映射的用户空间的起始地址，通常为 NULL（由内核来指定）
    // length：要映射的内存区域的大小
    // prot：期望的内存保护标志
    //      PROT_EXEC: 这个区域的页面由可以被 CPU 执行的指令组成
    //      PROT_READ: 这个区域的页面可读
    //      PROT_WRITE: 这个区域的页面的可写
    //      PROT_NONE: 这个区域内的页面不能被访问
    // flags：指定映射对象的类型
    //      MAP_ANON: 映射区不与任何文件关联（匿名），一般用于父子进程间通信
    //      MAP_PRIVATE: 被映射的对象是的私有、写时复制的
    //      MAP_SHARED: 被映射的对象是共享对象
    // fd：文件描述符（由open函数返回）
    // offset：设置在内核空间中已经分配好的的内存区域中的偏移，例如文件的偏移量，大小为PAGE_SIZE的整数倍
    // 返回值：mmap()返回被映射区的指针，该指针就是需要映射的内核空间在用户空间的虚拟地址
    char *start =
        mmap(NULL, sb.st_size, PROT_READ | PROT_WRITE, MAP_SHARED, fd, 0);
    if (start == MAP_FAILED) {
        fprintf(stderr, "open file error");
        return 0;
    }

    fprintf(stdout, start);

    return 0;
}
