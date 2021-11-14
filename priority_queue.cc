#include <iostream>
#include <queue>

using namespace std;

class Stu {
   private:
   public:
    int score;

    Stu(){};
    Stu(int score) { this->score = score; };
    ~Stu(){};
    bool operator<(const Stu& s) const { return this->score < s.score; };
    bool operator>(const Stu& s) const { return this->score > s.score; };
};

int main(int argc, char const* argv[]) {
    Stu s1(60);
    Stu s2(86);
    Stu s3(50);
    Stu s4(100);

    // 默认为大顶堆
    priority_queue<Stu> pq;
    pq.push(s1);
    pq.push(s2);
    pq.push(s3);
    pq.push(s4);
    while (pq.size() > 0) {
        cout << pq.top().score << endl;
        pq.pop();
    }

    cout << "================================" << endl;

    // 通过 greater 设置为小顶堆
    // 大顶堆对应的是 less<Stu>
    priority_queue<Stu, vector<Stu>, greater<Stu> > pq1;
    pq1.push(s1);
    pq1.push(s2);
    pq1.push(s3);
    pq1.push(s4);
    while (pq1.size() > 0) {
        cout << pq1.top().score << endl;
        pq1.pop();
    }

    return 0;
}
