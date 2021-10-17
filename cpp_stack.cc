#include <iostream>
#include <stack>

using namespace std;

int main(int argc, char const *argv[]) { 
    stack<int> s;
    s.push(1);
    auto t = s.top();
    cout << t << endl;
    return 0; 
}
