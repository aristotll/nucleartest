#include <iostream>
#include <string>

class Sales_data {
private:
    std::string bookNo_;
    unsigned units_gold_ = 0;
    double revenue_ = 0.0;
public:
    Sales_data(std::string bookNo, unsigned units_gold, double revenue);
    friend std::ostream& operator<<(std::ostream& io, Sales_data& obj); 
};

//
//int fn() {
//    struct Sales_data sale {
//        .bookNo = "123",
//        .units_gold = 123,
//        .revenue = 10.1,
//    };
//    Sales_data b, c;
//    std::cout << sale.bookNo << std::endl;
//}
