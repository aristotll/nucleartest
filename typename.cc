#include <iostream>
#include <unordered_map>

template <typename K, typename V>
void printMap(std::unordered_map<K, V>& map) {
    std::cout << "k: " << K << " v: " << V << std::endl;
}

int main() {
    std::unordered_map<std::string, std::string> map;
    map.emplace("name", "wangjunqiang");
    printMap(map);
}
