#include <iostream>
#include <algorithm>
#include <vector>

template <typename T>
std::ostream& operator<< (std::ostream& out, const std::vector<T>& vec) {
    if (!vec.empty()) {
        out << '[';
        std::copy(vec.begin(), vec.end(), std::ostream_iterator<T>(out, ", "));
        out << "\b\b]";
    }
    return out;
}

int main() {
    std::vector<int> vec{1, 2, 3};
    std::cout << vec << std::endl;
}
