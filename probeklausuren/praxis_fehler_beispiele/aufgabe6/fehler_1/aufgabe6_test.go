package aufgabe6

import "fmt"

func ExampleDuplicates() {
	l1 := []int{0, 10, 20, 30, 40, 10, 20, 50, 60, 10}
	l2 := []int{1, 2, 3, 4, 5, 4, 3, 2, 1}
	l3 := []int{1, 2, 3, 4}

	fmt.Println(Duplicates(l1))
	fmt.Println(Duplicates(l2))
	fmt.Println(Duplicates(l3))

	// Output:
	// [10 20]
	// [1 2 3 4]
	// []
}
