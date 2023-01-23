#include <sys/file.h>
#include <unistd.h>

#include <ctime>
#include <fstream>
#include <iostream>

// recordCurTimeToFile 写入 buf 和 当前时间到 fd
void recordCurTimeToFile(int fd, const char* buf, int bufsize) {
    time_t t = time(NULL);
    struct tm* lt = localtime(&t);
    char buffer[80];
    size_t buflen = strftime(buffer, 80, "%Y-%m-%d %H:%M:%S ", lt);
    write(fd, buffer, buflen);
    write(fd, buf, bufsize);
}

const char* curTime() {
    time_t t = time(NULL);
    struct tm* lt = localtime(&t);
    char buffer[80];
    size_t buflen = strftime(buffer, 80, "%Y-%m-%d %H:%M:%S ", lt);
    auto* str = new std::string(buffer, buflen);
    return str->c_str();
}

int main() {
    const char* filepath = "test.txt";
    // 运行前先清空文化内容
    open(filepath, O_TRUNC, 0777);
    pid_t pid = fork();
    int fd;

    if (pid == 0) {  // child
        fd = open(filepath, O_CREAT | O_RDWR | O_APPEND, 0777);
        //off_t offset = lseek(fd, 0, SEEK_CUR);
        //printf("%s [child] offset = %lld\n", curTime(), offset);
        printf("%s [child] open path = %s, fd = %d\n", curTime(), filepath, fd);
        printf("%s [child] try to lock\n", curTime());
        // 独占锁
        if ((flock(fd, LOCK_EX)) == -1) {
            perror("[child] lock file error");
            return 0;
        }
        printf("%s [child] locked\n", curTime());
        const char* writebuf = "[child] \n";
        recordCurTimeToFile(fd, writebuf, strlen(writebuf));
        sleep(3);
        // 释放锁
        if ((flock(fd, LOCK_UN)) == -1) {
            perror("[child] unlock file error");
            return 0;
        }
        printf("%s [child] release lock\n", curTime());
    } else if (pid > 0) {  // parent
        fd = open(filepath, O_CREAT | O_RDWR | O_APPEND, 0777);
        // int fd1 = dup(fd);
        printf("%s [parent] open path = %s, fd = %d\n", curTime(), filepath,
               fd);
        printf("%s [parent] try to lock\n", curTime());

        // 独占锁
        if ((flock(fd, LOCK_EX)) == -1) {
            perror("lock file error");
            return 0;
        }
        printf("%s [parent] locked\n", curTime());
        const char* writebuf = "[parent] \n";
        recordCurTimeToFile(fd, writebuf, strlen(writebuf));
        sleep(3);
        // 释放锁
        if ((flock(fd, LOCK_UN)) == -1) {
            perror("unlock file error");
            return 0;
        }
        printf("%s [parent] release lock\n", curTime());
        waitpid(pid, nullptr, 0);
    } else {
        perror("fork error");
        return 0;
    }
}
