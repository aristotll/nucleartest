#include <functional>
#include <iostream>
#include <vector>

enum Status { Running, Success, Failed };

// typedef Status std::string

class TODO {
   private:
    std::string name;
    Status status;

   public:
    Status getStatus();
    std::string getName();
    TODO(std::string name, Status status);
};

Status TODO::getStatus() { return this->status; }

std::string TODO::getName() { return this->name; }

TODO::TODO(std::string name, Status status) {
    this->name = name;
    this->status = status;
}

int main() {
    std::vector<std::function<bool(TODO)> > filters;
    filters.emplace_back([](TODO t) -> bool {
        if (t.getStatus() == Status::Running) {
            return true;
        }
        return false;
    });

    std::vector<TODO> vec1;
    vec1.emplace_back(TODO("event1", Status::Running));
    vec1.emplace_back(TODO("event2", Status::Success));

    for (auto &&i : vec1) {
        for (auto &&filter : filters) {
            if (filter(i)) {
                std::cout << i.getName() << std::endl;
            }
        }
    }
}
