package main

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

// func TestWithVariating(t *testing.T) {
// var eps float64 = 1e-5
// data := [17][4]float64{{103.33, 333.099, 9.9, 7},
// {2122.98, 202257.54, -4046, 7},
// {-96.73, -327.009, 9.9, 7},
// {-103.33, 333.099, -9.9, 7},
// {3347, 46711, 163317, 7.77},
// {110981, -2219920, 11100100, 7.77},
// {-101014, 404044, -404040, 1.111},
// {-99999.78, -21999.9879, -1210, 101},
// {-100007, 0, 0, 101},
// {11, 30.25, 0, 0.23},
// {-11, 30.25, 0, 0.23},
// {11.1, 41.07, 50.653, 0.23},
// {-21.3, 151.23, -357.911, 200},
// {0, 100.23, 0, 200},
// {-1.333, 1, -1.333, 1000},
// {9.999, 12345, 123437.655, 1000},
// {0, 0, 0, 1000}}
// }

func TestOneRoot(t *testing.T) {
	var eps float64 = 1e-5
	data := [][]float64{{11.1, 41.07, 50.653, 0.23},
		{-21.3, 151.23, -357.911, 200},
		{0, 100.23, 0, 200},
		{-1.333, 1, -1.333, 1000},
		{9.999, 12345, 123437.655, 1000},
		{0, 0, 0, 1000}}

	for _, d := range data {
		roots, values, _ := FindAnswer(d[0], d[1], d[2], eps, d[3])
		assert.Equalf(t, 1, len(values), "Expected result length 1, but got %v on values %v\n", len(values), d)
		assert.Equalf(t, 1, len(roots), "Expected result length 1, but got %v on roots %v\n", len(roots), d)
		for i, x := range values {
			assert.Truef(t, math.Abs(x) <= eps, "Expected -eps <= x <= eps (near 0), but got %v on i = %v with params %v\n", x, i, d)
		}

		for _, x := range roots {
			t.Logf("Root %f have a multiple of %v\n", x.Root, x.Multipl)
		}
		t.Logf("<==============================>\n")

		polynomial := func(x float64) float64 {
			return x*x*x + d[0]*x*x + d[1]*x + d[2]
		}

		for i, x := range roots {
			assert.Truef(t, math.Abs(polynomial(x.Root)) <= eps, "Expected -eps <= x <= eps (near 0), but got %v on i = %v with params %v\n", polynomial(x.Root), i, d)
		}
	}
}

func TestTwoRoots(t *testing.T) {
	var eps float64 = 1e-5
	data := [][]float64{{3347, 46711, 163317, 7.77},
		{110981, -2219920, 11100100, 7.77},
		{-101014, 404044, -404040, 1.111},
		{-99999.78, -21999.9879, -1210, 101},
		{-100007, 0, 0, 101},
		{11, 30.25, 0, 0.23},
		{-11, 30.25, 0, 0.23}}

	for _, d := range data {
		roots, values, _ := FindAnswer(d[0], d[1], d[2], eps, d[3])
		assert.Equalf(t, 2, len(values), "Expected result length 2, but got %v on values %v\n", len(values), d)
		assert.Equalf(t, 2, len(roots), "Expected result length 2, but got %v on roots %v\n", len(roots), d)
		for i, x := range values {
			assert.Truef(t, math.Abs(x) <= eps, "Expected -eps <= x <= eps (near 0), but got %v on i = %v with params %v\n", x, i, d)
		}

		for _, x := range roots {
			t.Logf("Root %f have a multiple of %v\n", x.Root, x.Multipl)
		}
		t.Logf("<==============================>\n")
		polynomial := func(x float64) float64 {
			return x*x*x + d[0]*x*x + d[1]*x + d[2]
		}

		for i, x := range roots {
			assert.Truef(t, math.Abs(polynomial(x.Root)) <= eps, "Expected -eps <= x <= eps (near 0), but got %v on i = %v with params %v\n", polynomial(x.Root), i, d)
		}

	}
}

func TestThreeRoots(t *testing.T) {
	var eps float64 = 1e-5
	data := [][]float64{{103.33, 333.099, 9.9, 7},
		{2122.98, 202257.54, -4046, 7},
		{-96.73, -327.009, 9.9, 7},
		{-103.33, 333.099, -9.9, 7}}

	for _, d := range data {
		roots, values, _ := FindAnswer(d[0], d[1], d[2], eps, d[3])
		assert.Equalf(t, 3, len(values), "Expected result length 3, but got %v on values %v\n", len(values), d)
		assert.Equalf(t, 3, len(roots), "Expected result length 3, but got %v on roots %v\n", len(roots), d)
		for i, x := range values {
			assert.Truef(t, math.Abs(x) <= eps, "Expected -eps <= x <= eps (near 0), but got %v on i = %v with params %v\n", x, i, d)
		}

		for _, x := range roots {
			t.Logf("Root %f have a multiple of %v\n", x.Root, x.Multipl)
		}
		t.Logf("<==============================>\n")

		polynomial := func(x float64) float64 {
			return x*x*x + d[0]*x*x + d[1]*x + d[2]
		}

		for i, x := range roots {
			assert.Truef(t, math.Abs(polynomial(x.Root)) <= eps, "Expected -eps <= x <= eps (near 0), but got %v on i = %v with params %v\n", polynomial(x.Root), i, d)
		}

	}
}

func TestVariatingEpsilon(t *testing.T) {
	a := -0.000001
	b := 1.
	c := -0.000001
	delta := 1000.
	eps := []float64{0.0000001, 0.0001, 0.1, 1}

	for _, e := range eps {
		roots, values, _ := FindAnswer(a, b, c, e, delta)
		assert.Equalf(t, 1, len(values), "Expected result length 1, but got %v on values %v with eps %v\n", len(values), []float64{a, b, c, delta}, e)
		assert.Equalf(t, 1, len(roots), "Expected result length 1, but got %v on roots %v with eps %v\n", len(roots), []float64{a, b, c, delta}, e)

		for i, x := range values {
			assert.Truef(t, math.Abs(x) <= e, "Expected -eps <= x <= eps (near 0), but got %v on i = %v with params %v and eps %v\n", x, i, []float64{a, b, c, delta}, e)
		}

		for _, x := range roots {
			t.Logf("Root %f have a multiple of %v\n", x.Root, x.Multipl)
		}
		t.Logf("<==============================>\n")

		polynomial := func(x float64) float64 {
			return x*x*x + a*x*x + b*x + c
		}

		for i, x := range roots {
			assert.Truef(t, math.Abs(polynomial(x.Root)) <= e, "Expected -eps <= x <= eps (near 0), but got %v on i = %v with params %v and eps %v\n", polynomial(x.Root), i, []float64{a, b, c, delta}, e)
		}
	}
}
