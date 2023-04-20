#include <iostream>
#include <string>

class Sales_data {
private:
    std::string bookNo_;
    unsigned units_sold_ = 0;
    double revenue_ = 0.0;
public:
    Sales_data(std::string bookNo, unsigned units_sold, double revenue) :
  bookNo_(bookNo), units_sold_(units_sold), revenue_(revenue)  {}

    friend std::ostream& operator<<(std::ostream& io, Sales_data& obj);
    friend std::istream& operator>>(std::istream& io, Sales_data& obj); 
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
