#include <iostream>

int main() {
  int *a[5];
  int *(&b)[5] = a;
}
