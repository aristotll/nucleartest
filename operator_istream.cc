#include <iostream>
#include "Sales_data.h"

std::istream &operator>>(std::istream &is, Sales_data &item) {
    double price; // 不需要初始化，因为我们将先读入数据到 price，然后才使用它
    is >> item.bookNo_ >> item.units_sold_ >> price;
    if (is) // 检查输入是否成功
        item.revenue_ = item.units_sold_ * price;
    else
        item = Sales_data("c++", 123, 456);    // 输入失败：对象被赋予默认的状态
}

int main() {
    Sales_data sale("c++", 123, 456);
    std::cin >> sale;
    std::cout << sale;
}
