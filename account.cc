#include <iostream>

class Account {
public:
    void calculate() { amount += amount * interestRate; }
    static double rate() { return interestRate; }
    static void rate(double);
private:
    std::string owner;
    double amount;
    static double interestRate; 
    static double initRate();
};

void Account::rate(double v) {
    interestRate = v;
}

double Account::interestRate = 100;

int main() {
    double r;
    r = Account::rate();
    std::cout << r << std::endl;
}
