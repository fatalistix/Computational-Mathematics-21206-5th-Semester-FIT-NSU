package main

import (
	"errors"
	"fmt"
	"math"
	"os"
)

type rootWithMultipl struct {
	Root    float64
	Multipl int
}

func findRootMonoInRightInfSemiInterval(polynomial func(x float64) float64, l, eps, delta float64, ascending bool) float64 {
	k := 1.
	if ascending {
		for polynomial(l+k*delta) < eps {
			k += 1.
		}
		return findRootMonoInLineSegment(polynomial, l+(k-1.)*delta, l+k*delta, eps)
	} else {
		for polynomial(l+k*delta) > -eps {
			k += 1.
		}
		return findRootMonoInLineSegment(polynomial, l+(k-1.)*delta, l+k*delta, eps)
	}
}

func findRootMonoInLeftInfSemiInterval(polynomial func(x float64) float64, r, eps, delta float64, ascending bool) float64 {
	k := 1.
	if ascending {
		for polynomial(r-k*delta) < eps {
			k += 1.
		}
		return findRootMonoInLineSegment(polynomial, r-k*delta, r-(k-1)*delta, eps)
	} else {
		for polynomial(r-k*delta) > -eps {
			k += 1.
		}
		return findRootMonoInLineSegment(polynomial, r-k*delta, r-(k-1)*delta, eps)
	}
}

func findRootMonoInLineSegment(polynomial func(x float64) float64, l, r, eps float64) float64 {
	if math.Abs(polynomial(l)) < eps {
		return l
	}
	if math.Abs(polynomial(r)) < eps {
		return r
	}

	c := (r + l) / 2.
	for math.Abs(polynomial(c)) >= eps {
		if polynomial(c) > eps {
			if polynomial(l) > eps {
				l = c
			} else {
				r = c
			}
		} else {
			if polynomial(l) > eps {
				r = c
			} else {
				l = c
			}
		}
		c = (r + l) / 2.
	}
	return c
}

// func truncate(x float64, precision int) float64 {
// 	precScaler := math.Pow10(precision)
// 	return math.Round(x*precScaler) / precScaler
// }

func printRootsWithValues(roots []rootWithMultipl, values []float64) {
	// buf1 := int(math.Floor(math.Log10(eps)))
	// buf2 := -int(math.Floor(math.Log10(eps)))
	// fmt.Println("PRECISION: ", precision, math.Log10(eps), math.Floor(math.Log10(eps)), buf1, buf2)

	if len(roots) == 1 {
		fmt.Println("У данного уравнения один вещественный корень")
		fmt.Println("Корень:   ", roots[0].Root, " кратности ", roots[0].Multipl)
		fmt.Println("Значение: ", values[0])
	} else {
		fmt.Printf("У данного уравнения %v вещественных корня\n", len(roots))
		fmt.Print("Корни:    ")
		for _, x := range roots {
			fmt.Print(x.Root, " кратности ", x.Multipl, "; ")
		}
		fmt.Println()
		fmt.Print("Значения: ")
		for _, x := range values {
			fmt.Print(x, "; ")
		}
		fmt.Println()
	}
}

func getRootsWithMultipleAndValues(polynomial, diffpolynomial func(x float64) float64, eps float64, roots ...float64) ([]rootWithMultipl, []float64) {
	// var precision int
	// if eps >= 1. {
	// precision = 0
	// } else {
	// precision = -int(math.Floor(math.Log10(eps)))
	// }

	if len(roots) == 1 {
		if math.Abs(diffpolynomial(roots[0])) <= eps {
			return []rootWithMultipl{{roots[0], 3}}, []float64{polynomial(roots[0])}
		}
		return []rootWithMultipl{{roots[0], 1}}, []float64{polynomial(roots[0])}
	} else {
		truncatedMultRoots := make([]rootWithMultipl, len(roots))
		truncatedValues := make([]float64, len(roots))

		for i, x := range roots {
			if math.Abs(diffpolynomial(x)) <= eps {
				truncatedMultRoots[i] = rootWithMultipl{x, 2}
			} else {
				truncatedMultRoots[i] = rootWithMultipl{x, 1}
			}
			truncatedValues[i] = polynomial(x)
		}

		return truncatedMultRoots, truncatedValues
	}
}

func correctInput(a, b, c, e, d float64) error {
	if e <= 0. {
		return errors.New("epsilon cannot be less then zero or be equal zero")
	}

	if d <= 0. {
		return errors.New("delta cannot be less then zero or be equal zero")
	}

	return nil
}

func FindAnswer(a, b, c, eps, delta float64) ([]rootWithMultipl, []float64, error) {
	err := correctInput(a, b, c, eps, delta)
	if err != nil {
		return nil, nil, err
	}

	var polynomial func(x float64) float64 = func(x float64) float64 {
		return x*x*x + a*x*x + b*x + c
	}

	var diffpolynomial func(x float64) float64 = func(x float64) float64 {
		return 3*x*x + 2*a*x + b
	}

	// discr := b*b - 4.*a*c
	discr := 4.*a*a - 4.*b*3.
	// fmt.Println("DISCR", discr)

	if discr <= 10e-16 {
		// fmt.Println("DISCR IS <= 10e-16!!!")
		valAtZero := polynomial(0.)
		if math.Abs(valAtZero) < eps {
			// fmt.Println("FIRST AFT DISCR")
			roots, values := getRootsWithMultipleAndValues(polynomial, diffpolynomial, eps, 0.)
			return roots, values, nil
		} else if valAtZero < -eps {
			// fmt.Println("SECOND AFT DISCR")
			roots, values := getRootsWithMultipleAndValues(polynomial, diffpolynomial, eps, findRootMonoInRightInfSemiInterval(polynomial, 0., eps, delta, true))
			return roots, values, nil
		} else {
			// fmt.Println("THIRD AFT DISCR")
			roots, values := getRootsWithMultipleAndValues(polynomial, diffpolynomial, eps, findRootMonoInLeftInfSemiInterval(polynomial, 0., eps, delta, false))
			return roots, values, nil
		}
	}

	x1 := (-2*a + math.Sqrt(discr)) / (2. * 3.)
	x2 := (-2*a - math.Sqrt(discr)) / (2. * 3.)

	alpha := math.Min(x1, x2)
	beta := math.Max(x1, x2)

	palpha := polynomial(alpha)
	pbeta := polynomial(beta)

	if palpha < -eps && pbeta < -eps {
		// fmt.Println("FIRST")
		roots, values := getRootsWithMultipleAndValues(polynomial, diffpolynomial, eps, findRootMonoInRightInfSemiInterval(polynomial, beta, eps, delta, true))
		return roots, values, nil
	} else if math.Abs(palpha) <= eps && pbeta < -eps {
		// fmt.Println("SECOND")
		roots, values := getRootsWithMultipleAndValues(polynomial, diffpolynomial, eps, alpha, findRootMonoInRightInfSemiInterval(polynomial, beta, eps, delta, true))
		return roots, values, nil
	} else if palpha > eps && pbeta < -eps {
		// fmt.Println("THIRD")
		roots, values := getRootsWithMultipleAndValues(polynomial, diffpolynomial, eps, findRootMonoInLeftInfSemiInterval(polynomial, alpha, eps, delta, false), findRootMonoInLineSegment(polynomial, alpha, beta, eps), findRootMonoInRightInfSemiInterval(polynomial, beta, eps, delta, true))
		return roots, values, nil
	} else if palpha > eps && math.Abs(pbeta) <= eps {
		// fmt.Println("FORTH")
		roots, values := getRootsWithMultipleAndValues(polynomial, diffpolynomial, eps, findRootMonoInLeftInfSemiInterval(polynomial, alpha, eps, delta, false), beta)
		return roots, values, nil
	} else if palpha > eps && pbeta > eps {
		// fmt.Println("FIFTH")
		roots, values := getRootsWithMultipleAndValues(polynomial, diffpolynomial, eps, findRootMonoInLeftInfSemiInterval(polynomial, alpha, eps, delta, false))
		return roots, values, nil
	} else {
		// fmt.Println("SIXTH")
		roots, values := getRootsWithMultipleAndValues(polynomial, diffpolynomial, eps, (alpha+beta)/2.)
		return roots, values, nil
	}
}

func main() {
	var (
		a     float64
		b     float64
		c     float64
		eps   float64
		delta float64
		// precision int
	)

	fmt.Println("Нахождение корней кубического уравнения с точностью до epsilon")
	fmt.Println("Решение уравнения формата x^3 + a*x^2 + b*x + c")
	fmt.Println("Параметр delta отвечает за перемещение промежутка на полуинтервалах")
	// fmt.Println("Параметр precision отвечает за количество знаков после запятой (0 <= precision <= 15)")
	fmt.Println("Введите коэффициенты в следующем порядке: a, b, c, eps, delta")

	// _, err := fmt.Scan(&a, &b, &c, &eps, &delta, &precision)
	// if err != nil {
	// fmt.Fprintln(os.Stderr, err)
	// return
	// }

	_, err := fmt.Scan(&a)
	if err != nil {
		fmt.Fprintln(os.Stderr, "a должно быть дробным числом, записанным через точку!")
		return
	}

	_, err = fmt.Scan(&b)
	if err != nil {
		fmt.Fprintln(os.Stderr, "b должно быть дробным числом, записанным через точку!")
		return
	}

	_, err = fmt.Scan(&c)
	if err != nil {
		fmt.Fprintln(os.Stderr, "c должно быть дробным числом, записанным через точку!")
		return
	}

	_, err = fmt.Scan(&eps)
	if err != nil {
		fmt.Fprintln(os.Stderr, "eps должно быть дробным числом, записанным через точку!")
		return
	}

	_, err = fmt.Scan(&delta)
	if err != nil {
		fmt.Fprintln(os.Stderr, "delta должно быть дробным числом, записанным через точку!")
		return
	}

	roots, values, err := FindAnswer(a, b, c, eps, delta)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	// fmt.Println("LENGTH", len(roots), len(values))

	printRootsWithValues(roots, values)

	// _, err = fmt.Scan(&precision)
	// if err != nil {
	// fmt.Fprintln(os.Stderr, "precision должно быть целым числом!")
	// return
	// }

}

// fmt.Println(findRootMonoInRightInfSemiInterval(func(x float64) float64 {
// return x*x - 99
// }, 1, 0.001, 3, true))
