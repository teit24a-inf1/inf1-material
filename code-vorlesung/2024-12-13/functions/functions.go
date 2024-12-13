package functions

import "math"

// F: Z -> Z
// F(x) = x^2
func F(x int) int {
	return x * x
}

// G: R -> R
func G(x float64) float64 {
	return math.Pow(x, 2)
	//return x * x
}
