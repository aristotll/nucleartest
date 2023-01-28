#include <iostream>
#include <string>

struct Sales_data {
    std::string bookNo;
    unsigned units_gold = 0;
    double revenue = 0.0;
};

int main() {
    struct Sales_data sale {
        .bookNo = "123",
        .units_gold = 123,
        .revenue = 10.1,
    };
    Sales_data b, c;
    std::cout << sale.bookNo << std::endl;
}
