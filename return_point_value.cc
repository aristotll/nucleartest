#include <iostream>
#include <vector>

using namespace std;

class TreeNode {
   private:
   public:
    int value;
    TreeNode *left;
    TreeNode *right;

    TreeNode();
    TreeNode(int value, TreeNode *left, TreeNode *right);
    ~TreeNode();
};

TreeNode::TreeNode() {}

TreeNode::TreeNode(int value, TreeNode *left, TreeNode *right) {
    this->left = left;
    this->right = right;
    this->value = value;
}

TreeNode::~TreeNode() {}

vector<TreeNode *> inorder() {
    vector<TreeNode *> v;
    TreeNode *t1 = new TreeNode(1, nullptr, nullptr);
    TreeNode *t2 = new TreeNode(2, nullptr, nullptr);
    v.emplace_back(t1);
    v.emplace_back(t2);
    return v;
}

int main(int argc, char const *argv[]) {
    auto vc = inorder();
    auto val = vc[0]->value;
    cout << val << endl;
    return 0;
}

int main(int argc, char const *argv[]) {
    /* code */
    return 0;
}

