#include <iostream>

using namespace std;

namespace firstSpace {
int a = 10;
void func() { cout << "func1" << endl; }
} // namespace firstSpace

namespace secondSpace {
int a = 20;
void func() { cout << "func2" << endl; }
} // namespace secondSpace

class Diff {
private:
  int x;

public:
  Diff();
  Diff(int x);
  ~Diff(){};
  void func();
  void printX();
};

Diff::Diff() {}
Diff::Diff(int x) { this->x = x; }
void Diff::func() { cout << "func" << endl; }
void Diff::printX() { cout << this->x << endl; }

int main(int argc, char const *argv[]) {
  firstSpace::func();
  secondSpace::func();

  cout << firstSpace::a << endl;
  cout << secondSpace::a << endl;

  Diff d;
  d.func();

  Diff d1(1);

  return 0;
}
