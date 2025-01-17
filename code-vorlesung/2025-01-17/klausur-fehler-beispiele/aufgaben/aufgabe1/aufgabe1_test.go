package aufgabe1

import (
	"fmt"
)

func ExampleEvenPrefix() {
	l1 := []int{1, 2, 3, 4, 5}
	l2 := []int{2, 3, 4, 5, 6}
	l3 := []int{2, 4, 6, 7, 8}
	l4 := []int{}

	fmt.Println(EvenPrefix(l1))
	fmt.Println(EvenPrefix(l2))
	fmt.Println(EvenPrefix(l3))
	fmt.Println(EvenPrefix(l4))

	// Output:
	// []
	// [2]
	// [2 4 6]
	// []
}
