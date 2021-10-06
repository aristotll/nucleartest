#include <iostream>

using namespace std;

class Animal {
private:
    /* data */
public:
    Animal(/* args */) {}
    ~Animal() {cout << "[Animal] is delete" << endl;}

    string name;

    virtual void eat() {
        cout << "[father virtual] eat" << endl;
    }
    void run() {
        cout << "[father] run" << endl;
    }
};

class Dog : public Animal {
private:
    /* data */
public:
    Dog(/* args */) {}
    ~Dog() {cout << "[Dog] is delete" << endl;}

};

class Dog1 : public Animal {
private:
    /* data */
public:
    Dog1(/* args */) {}
    ~Dog1() {cout << "[Dog1] is delete" << endl;}

    void eat() {
        cout << "[child] eat" << endl;
    }

    void run() {
        cout << "[child] run" << endl;
    }
};

int main(int argc, char const *argv[])
{
    /* code */
    Animal *dog = new Dog();
    dog->name = "dog";
    dog->eat();
    dog->run();
    cout << dog->name << endl;
    delete(dog);

    cout << "================" << endl;

    Animal *dog1 = new Dog1();
    dog1->eat();
    dog1->run();
    delete(dog1);

    return 0;
}
