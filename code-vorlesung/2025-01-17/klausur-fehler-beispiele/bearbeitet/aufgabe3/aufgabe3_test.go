package aufgabe3

import "fmt"

func ExampleElementDiffs() {
	l1 := []int{1, 2, 1, 2, 1}
	l2 := []int{3, 3, 3, 3, 3}
	l3 := []int{2, 2, 2}

	fmt.Println(ElementDiffs(l1, l2))
	fmt.Println(ElementDiffs(l2, l1))
	fmt.Println(ElementDiffs(l1, l3))

	// Output:
	// [-2 -1 -2 -1 -2]
	// [2 1 2 1 2]
	// [-1 0 -1 2 1]
}
