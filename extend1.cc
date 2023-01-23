#include <iostream>

class Father {
   protected:
    int x;

   public:
    Father();
    ~Father();
    friend std::ostream& operator<<(std::ostream& os, const Father& obj);
};

Father::Father() { this->x = 1; }

Father::~Father() {}

class Child : Father {
   private:
    int y_;

   public:
    Child() = default;
    Child(int y);
    friend std::ostream& operator<<(std::ostream& os, const Child& obj);
};

Child::Child(int y) { this->y_ = y; }

std::ostream& operator<<(std::ostream& os, const Child& obj) {
    os << "x: " << obj.x << " y: " << obj.y_;
    return os;
}

std::ostream& operator<<(std::ostream& os, const Father& obj) {
    os << obj.x;
    return os;
}

int main() {
    Child child;
    std::cout << child << std::endl;

    Father father;
    std::cout << father << std::endl;
}
