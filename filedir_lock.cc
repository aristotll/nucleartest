#include <dirent.h>
#include <sys/file.h>

#include <ctime>
#include <fstream>
#include <iostream>

class FileMutex {
   private:
    std::filesystem::path filepath;
    std::string pattern;
    int fd;
    bool printLog;

   public:
    FileMutex(std::filesystem::path filepath);
    FileMutex(std::filesystem::path filepath, std::string pattern);
    FileMutex(std::filesystem::path filepath, std::string pattern,
              bool printLog);
    ~FileMutex();
    void open();
    void lock();
    void unlock();
    void rlock();
    void runlock();
    void close();
    void tryLock();
};

FileMutex::FileMutex(std::filesystem::path filepath) {
    // this->pattern = pattern;
    // // 创建临时文件夹
    // this->temp_dir = std::filesystem::temp_directory_path();
    // this->temp_dir /= this->pattern;
    // std::cout << this->temp_dir << std::endl;
    // std::filesystem::create_directory(this->temp_dir);
    this->filepath = filepath;
}

FileMutex::FileMutex(std::filesystem::path filepath, std::string pattern)
    : FileMutex(filepath) {
    // this->filepath = filepath;
    this->pattern = pattern;
}

FileMutex::FileMutex(std::filesystem::path filepath, std::string pattern,
                     bool printLog)
    : FileMutex(filepath, pattern) {
    this->printLog = printLog;
    // Tmp(pattern);
}

FileMutex::~FileMutex() { std::filesystem::remove_all(this->temp_dir); }

void FileMutex::open() {
    if (std::filesystem::is_regular_file(this->path)) {
        auto dir = opendir(this->filepath.c_str());
        this->fd = dirfd(dir);

        if (this->fd == -1) {
            throw std::runtime_error("open dir error");
            // perror("open file error");
            // return;
        }
    } else {
        this->fd = opendir(this->filepath.c_str(), O_CREAT | O_RDWR);
        if (this->fd == -1) {
            throw std::runtime_error("open file error");
        }
    }

    if (this->printLog) {
        std::cout << "path: " << this->filepath.c_str() << " fd: " << fd
                  << std::endl;
    }
}

int main() {
    auto path = std::filesystem::temp_directory_path();

    auto FileMutex = FileMutex("test123", true);
    try {
        FileMutex.open();
    } catch (const std::exception& e) {
        std::cerr << e.what() << '\n';
    }
}
