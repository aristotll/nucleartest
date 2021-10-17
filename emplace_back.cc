#include <iostream>
#include <vector>

using namespace std;

#define printVector(x) for (auto v : x) {cout << v << " ";}

int main(int argc, char const *argv[]) {
    vector<int> vc;
    vc.emplace_back(1);
    vc.emplace_back(2);

    printVector(vc);
    cout << endl;
    return 0;
}
