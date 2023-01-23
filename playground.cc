#include <iostream>

int txt_size() {
    return 1;
}

int main() {
    unsigned buf_size = 1024;
    int ia[buf_size];
    int ia1[4 * 7 - 14];
    int ia2[txt_size()];
    char st[11] = "fundamental";
}
