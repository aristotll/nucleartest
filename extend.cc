#include <iostream>

class Quote {
public:
    Quote() = default;

    std::string isbn() const;
    virtual double net_price(std::size_t n) const;
};

double Quote::net_price(std::size_t n) const {
    return 3.14 * n;
}

std::string Quote::isbn() const {
    return "true";
}

class Bulk_quote : public Quote {
public:
    virtual double net_price(std::size_t n) const;
};

double Bulk_quote::net_price(std::size_t n) const {
    return 3.14 * n;
}

double print_total(std::ostream &os, const Quote &item, size_t n) {
    double ret = item.net_price(n);
    os << "ISBN: " << item.isbn() << " # sold: " << n << " total due: " << ret << std::endl;
    return ret;
}

int main() {
    auto quote = new Quote();
    print_total(std::cout, *quote, 10);

    auto quote1 = new Bulk_quote();
    print_total(std::cout, *quote1, 20);
}
