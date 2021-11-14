#include <cstdlib>
#include <iostream>

using namespace std;

int main(int argc, char const *argv[], char *envp[]) {
    cout << "Command-ine arguments: " << endl;
    for (int i = 0; i < argc; i++) {
        cout << "\targv[" << i << "]: " << argv[i] << endl;
    }

    cout << endl;
    cout << "Environment Variables: " << endl;
    for (int i = 0; ;i++) {
        if (envp[i] == nullptr) {
            break;
        }
        cout << "\tenv[" << i << "]: " << envp[i] << endl;
    }

    return 0;
}
