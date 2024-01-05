/*
 * Copyright (c) 2023, Balashov Vyacheslav
 */

#include "./numerical-integration.h"
#include <cmath>
#include <iomanip>
#include <iostream>
#include <vector>

int main() {
  double a = 0.;
  double b = 1.;
  int n1 = 10;
  int n2 = 20;
  double h1 = (b - a) / n1;
  double h2 = (b - a) / n2;

  std::vector<double> v1(n1 + 1);
  std::vector<double> v2(n2 + 1);

  for (int i = 0; i < v1.size(); ++i) {
    v1[i] = std::exp(a + i * h1);
  }

  for (int i = 0; i < v2.size(); ++i) {
    v2[i] = std::exp(a + i * h2);
  }

  std::cout << std::fixed << std::setprecision(10)
            << numericalTrapezoidIntegration(h1, v1) << ' '
            << numericalTrapezoidIntegration(h2, v2) << std::endl;

  return 0;
}
