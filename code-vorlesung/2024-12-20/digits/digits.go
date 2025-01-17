package digits

import (
	"slices"
)

// Erwartet eine Zahl n und liefert
// eine Liste der Ziffern von n.
func Digits(n int) []int {
	result := []int{}

	for n != 0 {
		r := n % 10
		n = n / 10

		result = append(result, r)
	}

	slices.Reverse(result)
	return result
}
