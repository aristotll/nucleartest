#include <iostream>

class Foo {
private:
  int val;
  Foo(int i) : val(i) {}
};

class Bar : public Foo {
public:
  char *str;
  int i;
  Bar(int i, char *s) {
    i = i;
    str = s;
  }
};

int main() {
    
}
