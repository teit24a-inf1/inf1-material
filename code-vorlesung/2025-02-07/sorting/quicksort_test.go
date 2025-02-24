package sorting

import "fmt"

func ExampleQuickSort() {
	l1 := []int{25, 3, 2, 17, 32, 38}

	QuickSort(l1)

	fmt.Println(l1)

	// Output:
	// [2 3 17 25 32 38]
}
