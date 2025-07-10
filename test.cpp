#include <iostream>

using namespace std;

int main() {
    cout << "Hello, World!" << endl;

    int a;
    int b;
    cin >> a;
    cin >> b;
    int c = a + b;
    cout << "The sum of " << a << \
            " and " << b << \
            " is " << c << "." << endl;
    return 0;
}