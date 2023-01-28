#include <iostream>
#include <ranges>
#include <vector>

int main() {
    std::vector<int> v{1, 2, 3, 4, 5};
    for (auto i : std::ranges::input_range<std::vector<int>>(v)) {
        std::cout << i << ' ';
    }
}

