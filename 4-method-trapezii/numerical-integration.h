/*
 * Copyright (c) 2023, Vyacheslav Balashov
 */

#pragma once

#include <cstdint>
#include <vector>

double numericalTrapezoidIntegration(const double h,
                                     const std::vector<double> &f);
