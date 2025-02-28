package aufgabe4

import "fmt"

func ExampleElementProducts() {
	l1 := []int{1, 2, 1, 2, 1}
	l2 := []int{3, 3, 3, 3, 3}
	l3 := []int{2, 2, 2}

	fmt.Println(ElementProducts(l1, l2))
	fmt.Println(ElementProducts(l2, l1))
	fmt.Println(ElementProducts(l1, l3))

	// Output:
	// [3 6 3 6 3]
	// [3 6 3 6 3]
	// [2 4 2 2 1]
}
