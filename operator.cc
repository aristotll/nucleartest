#include <iostream>

using namespace std;

class Student {
   private:
   public:
    string name;
    int age;

    Student();
    Student(string name, int age);
    ~Student();
    // 重载 + 号
    Student operator+(const Student& stu) {
        Student s;
        s.name = this->name + stu.name;
        s.age = this->age + stu.age;
        return s;
    }

    // 重载 > 号
    bool operator>(const Student& stu) { return this->age > stu.age; }
};

Student::Student() {}

Student::Student(string name, int age) {
    this->name = name;
    this->age = age;
}

Student::~Student() {}

int main(int argc, char const* argv[]) {
    Student s1("zhang3", 12);
    Student s2("li4", 20);

    Student ss = s1 + s2;

    cout << ss.name << "  " << ss.age << endl;

    s1 > s2 ? cout << "s1 > s2" << endl : cout << "s1 < s2" << endl;

    return 0;
}
