#include <iostream>
#include <thread>
#include <mutex>

int v = 1;

// 临界区
void critical_section() {
    static std::mutex mutex;
    std::lock_guard<std::mutex> lock(mutex);

}

int main() {
    
}
