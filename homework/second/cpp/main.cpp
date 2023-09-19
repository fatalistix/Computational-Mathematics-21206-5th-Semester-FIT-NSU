/*
 * Copyright (c) 2023, Balashov Vyacheslav
 */

#include <cmath>
#include <iostream>

int main() {
  auto foo = [](double x) { return x * x * x - 10 * x - 1; };

  auto diffoo = [](double x) { return (x * x * x - 1) / 10; };

  auto revfoo = [](double x) { return std::pow(10 * x + 1, 1. / 3.); };

  double x;
  while (true) {
    std::cout << "Insert x: " << std::flush;
    std::cin >> x;
    std::cout << "foo: " << foo(x) << "; diffoo: " << diffoo(x)
              << "; refoo: " << revfoo(x) << std::endl;
  }
  return 0;
}
