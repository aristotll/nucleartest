#include "Sales_data.h"

std::ostream& operator<<(std::ostream& io, Sales_data& obj) {
    io << "bookNo: " << obj.bookNo_ << 
        " units_sold: " << obj.units_gold_ << 
        " revenue: " << obj.revenue_;
    return io; 
}

Sales_data::Sales_data(std::string bookNo, unsigned units_gold, double revenue) :
  bookNo_(bookNo), units_gold_(units_gold), revenue_(revenue)  {}

int main() {
    Sales_data data1("wangjunqiang", 1, 100);
    std::cout << data1 << std::endl;
}
