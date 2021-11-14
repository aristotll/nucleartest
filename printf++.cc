#include <iostream>
#include <vector>

using namespace std;

int main(int argc, char const *argv[]) {
    vector<int> vec{1, 2, 3, 4, 5};
    int i = 0;

    while (true) {
        cout << vec[i++] << endl;
        if (i == vec.size()) {
            break;
        }
    }

    return 0;
}
