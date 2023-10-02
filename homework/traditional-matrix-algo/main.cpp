/*
 * Copyright (c) 2023, Vyacheslav Balashov
 */

#include <iostream>
#include <vector>

using namespace std;

int main() {
  int n;
  cout << "Please enter matrix's size: " << flush;
  cin >> n;
  cout << "Please enter matrix:" << endl;
  vector<double> a(n), b(n), c(n), f(n), x(n);
  double value;
  for (int i = 0; i < n; ++i) {
    for (int j = 0; j < n; ++j) {
      cin >> value;
      if (i == j) {
        c[i] = value;
      } else if (i == j + 1) {
        a[i] = -value;
      } else if (i + 1 == j) {
        b[i] = -value;
      }
    }
  }

  cout << "Please enter result vector:" << endl;

  for (int i = 0; i < n; ++i) {
    cin >> value;
    f[i] = value;
  }

  vector<double> alpha(n), beta(n);
  alpha[0] = b[0] / c[0];
  beta[0] = f[0] / c[0];

  for (int i = 1; i < n; ++i) {
    alpha[i] = b[i] / (c[i] - a[i] * alpha[i - 1]);
    beta[i] = (f[i] + a[i] * beta[i - 1]) / (c[i] - a[i] * alpha[i - 1]);
  }

  x[n - 1] = (f[n - 1] + a[n - 1] * beta[n - 2]) /
             (c[n - 1] - a[n - 1] * alpha[n - 2]);
  for (int i = n - 2; i >= 0; --i) {
    x[i] = alpha[i] * x[i + 1] + beta[i];
  }

  cout << "Answer:\nx = [";
  for (int i = 0; i < n - 1; ++i) {
    cout << x[i] << ", ";
  }
  cout << x[n - 1] << "]" << endl;
  return 0;
}
