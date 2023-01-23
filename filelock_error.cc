// #include <strings.h>
#include <sys/file.h>
#include <unistd.h>

#include <ctime>
#include <fstream>
#include <iostream>

#define __O_RDWR__ std::ios_base::out | std::ios_base::in
#define __O_APPEND__ std::ios_base::app
#define __O_TRUNC__ std::ios_base::trunc

int main() {
    std::ofstream outfile;
    try {
        outfile.open("test.txt", __O_RDWR__ | __O_APPEND__ | __O_TRUNC__);
    } catch (std::ios_base::failure &e) {
        std::cerr << "Exception opening/reading/closing file\n";
    }

    auto helper = [](std::filebuf &fb) -> int {
        class Helper : public std::filebuf {
           public:
            int handle() { return _M_file.fd(); }
        };
        return static_cast<Helper &>(fb).handle();
    };

    int fd = helper(*outfile.rdbuf());

    // int fd;
    // char *pathname = "test.txt";
    // fd = open(pathname, O_CREAT | O_RDWR | O_TRUNC, 0777);
    // if (fd < 0) {
    //     fprintf(stderr, "open file error");
    //     return 0;
    // }

    int pid = fork();
    switch (pid) {
        case 0:  // child
        {
            // 加独占锁
            int lock = flock(fd, LOCK_EX);
            if (lock == -1) {
                perror("lock file error");
                return 0;
            }
            std::time_t now = std::chrono::system_clock::to_time_t(
                std::chrono::system_clock::now());
            outfile << "child process content written at: " << std::ctime(&now)
                    << std::endl;
            outfile.close();

            sleep(10);
            // 释放锁
            flock(fd, LOCK_UN);
            break;
        }
        case 1:  // parent
        {
            // 加独占锁
            int lock = flock(fd, LOCK_EX);
            if (lock == -1) {
                perror("lock file error");
                return 0;
            }
            std::time_t now = std::chrono::system_clock::to_time_t(
                std::chrono::system_clock::now());
            outfile << "parent process content written at: " << std::ctime(&now)
                    << std::endl;
            outfile.close();
            break;
        }
        default:  // error
        {
            perror("Error occurred while calling fork()\n");
            break;
        }
    }
}
