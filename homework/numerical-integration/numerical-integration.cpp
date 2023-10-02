/*
 * Copyright (c) 2023, Vyacheslav Balashov
 */

#include "./numerical-integration.h"

double numericalTrapezoidIntegration(const double h,
                                     const std::vector<double> &f) {
  double result = 0;
  for (int i = 0; i < f.size() - 1; ++i) {
    result += (f[i] + f[i + 1]) / 2 * h;
  }
  return result;
}
