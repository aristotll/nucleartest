#include <iostream>

using namespace std;

long exchange(long *xp, long y) {
    long x = *xp;
    *xp = y;
    return x;
}

int main() {
    long xp = 10;
    long y = 20;

    long ret = exchange(&xp, y);

    cout << "xp: " << xp << ", "
         << "y: " << y << ","
         "ret: " << ret << endl;
}