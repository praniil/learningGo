package numericalmethods

import (
	"fmt"
	"math"
)

func f(x float64) float64 {
	return x*x - 4
}

func solveBisec(a float64, b float64, tol float64, maxIter int) (float64, error) {
	if f(a)*f(b) >= 0 {
		return 0, fmt.Errorf("no root found in the given interval")
	}

	var root float64
	for i := 1; i <= maxIter; i++ {
		root = (a + b) / 2

		if math.Abs(f(root)) < tol {
			break
		}

		if f(a)*f(root) < 0 {
			b = root
		} else {
			a = root
		}
	}
	return root, nil
}

func Bisection() {
	fmt.Println("Bisection")
	a := 0.0
	b := 5.0

	tolerance := 0.00001
	maxIteration := 100
	root, err := solveBisec(a, b, tolerance, maxIteration)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("the root is %0.6f", root)
	}
}
