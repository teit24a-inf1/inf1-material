package functions

import (
	"fmt"
	"slices"
)

func Example_sort() {
	l1 := []int{17, 2, 42, 25, 38}
	slices.Sort(l1)
	fmt.Println(l1)

	l2 := []int{17, 2, 42, 25, 38}
	slices.SortFunc(l2, func(a, b int) int {
		return b - a
	})
	fmt.Println(l2)

	// Output:
}
