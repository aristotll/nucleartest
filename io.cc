#include <iostream>

int main() {
    // std::istream input (NULL);
    std::string str;
    // input >> str;
    for (;;) {
        //auto state = std::cin.rdstate();
        if (std::cin.eof()) {
            std::cout << "EOF!" << std::endl;
            break;
        }
        std::cin >> str;
        std::cout << str << std::endl;
    }
    // std::ofstream out1, out2;
}
