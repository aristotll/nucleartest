#include <iostream>

using namespace std;

class TODO {
private:
    int x_;
    string y_;

public:    
    TODO() {
        this->x_ = 100;
        this->y_ = "abc";
    }

    TODO(int x, string y) : x_(x), y_(y) {}

    friend ostream& operator<<(ostream &os, const TODO &t);
};

ostream& operator<<(ostream &os, const TODO &t) {
    os << "x: " << t.x_ << " y: " << t.y_;
    return os;
}

int main() {
    TODO todo;
    cout << todo << endl;
}
