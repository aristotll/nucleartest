#include <iostream>

void print(int &arr[5]) {
    for (auto elem : arr) {
        std::cout << elem << std::endl;
    }
}

int main() {
    int a[5] = {1, 2, 3, 4, 5};
    print(a);

    int &arr[5] = a;
}
