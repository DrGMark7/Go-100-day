#include <iostream>
#include <vector>

using namespace std;

int main() {
    vector<int> x(5);
    vector<int> h(5);

    for (int i = 0; i < 5; ++i) {
        x[i] = i + 1; // Example values for x
        h[i] = i + 2; // Example values for h
    }

    vector<int> y(x.size());

    for (size_t i = 0; i < x.size(); ++i) {
        y[i] = 0;
        for (size_t j = 0; j < h.size(); ++j) {
            y[i] += x[i] * h[j];
        }
    }
    for (int val : y) {
        cout << val << " ";
    }
    cout << endl;
}