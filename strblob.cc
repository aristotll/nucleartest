#include <iostream>
#include <vector>

class StrBlob {
   public:
    typedef std::vector<std::string>::size_type size_type;
    StrBlob();
    StrBlob(std::initializer_list<std::string> il);
    size_type size() const;
    bool empty() const;
    void push_back(const std::string &t);
    void pop_back();
    std::string &front();
    std::string &back();
    friend std::ostream &operator<<(std::ostream &out, const StrBlob &sb);

   private:
    std::shared_ptr<std::vector<std::string>> data;
    void check(size_type i, const std::string &msg) const;
};

std::ostream &operator<<(std::ostream &out, const StrBlob &sb) {
    if (!sb.empty()) {
        out << '[';
        std::copy(sb.data->begin(), sb.data->end(),
                  std::ostream_iterator<std::string>(out, ", "));
        out << "\b\b]";
    }
    return out;
}

StrBlob::StrBlob() : data(std::make_shared<std::vector<std::string>>()) {}
StrBlob::StrBlob(std::initializer_list<std::string> il)
    : data(std::make_shared<std::vector<std::string>>(il)) {}

StrBlob::size_type StrBlob::size() const { return this->data->size(); }

bool StrBlob::empty() const { return this->data->empty(); }
void StrBlob::check(size_type i, const std::string &msg) const {
    if (i >= this->data->size()) {
        throw std::out_of_range(msg);
    }
}

std::string &StrBlob::front() { return this->data->front(); }
std::string &StrBlob::back() { return this->data->back(); }

void StrBlob::push_back(const std::string &t) {
    this->check(0, "front on empty StrBlob");
    this->data->push_back(t);
}

void StrBlob::pop_back() {
    this->check(0, "back on empty StrBlob");
    this->data->pop_back();
}

int main() {
    StrBlob b1;
    {
        StrBlob b2 = {"a", "an", "the"};
        b1 = b2;
        b2.push_back("about");
        std::cout << "in scope, b1: " << b1 << " b2: " << b2 << std::endl;
    }
    std::cout << "out of scope, b1: " << b1 << std::endl;
}