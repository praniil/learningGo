//based on the idea of using the tangent line to approximate the root of a function

package numericalmethods

import (
	"fmt"
	"math"
)

func operatingNewtonRaphson(f func(float64) float64, fderi func(float64) float64, initialGuess float64, epsilon float64) float64 {
	x := initialGuess

	for math.Abs(f(x)) > epsilon {
		x = x - f(x)/fderi(x)
	}
	return x
}

func NewtonRaphson() {
	f := func(number float64) float64 {
		return number*number - 64
	}

	fderi := func(number float64) float64 {
		return 2 * number
	}
	initialGuess := 10.00
	epsilon := 0.00000
	result := operatingNewtonRaphson(f, fderi, initialGuess, epsilon)
	fmt.Println("result: ", result)
}
