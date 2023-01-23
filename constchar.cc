#include <iostream>

void print(const char* cp) {
    *cp = "123";
    if (cp) {
        while (*cp)
            std::cout << *cp++;
    }
}

int main() {
    const char* cp = "abc";
    print(cp);
}
