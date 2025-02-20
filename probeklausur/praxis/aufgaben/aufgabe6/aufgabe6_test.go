package aufgabe6

import "fmt"

func ExampleSymmetricDifference_standard_cases() {
	l1 := []int{20, 0, 15, 5, 10}
	l2 := []int{0, 10, 20, 30, 40}
	l3 := []int{15, 25, 5, -5, 35, 5, 25}

	fmt.Println(SymmetricDifference(l1, l2))
	fmt.Println(SymmetricDifference(l1, l3))
	fmt.Println(SymmetricDifference(l2, l3))

	// Output:
	// [15 5 30 40]
	// [20 0 10 25 -5 35 25]
	// [0 10 20 30 40 15 25 5 -5 35 5 25]
}

func ExampleSymmetricDifference_reversed_cases() {
	l1 := []int{20, 0, 15, 5, 10}
	l2 := []int{0, 10, 20, 30, 40}
	l3 := []int{15, 25, 5, -5, 35, 5, 25}

	fmt.Println(SymmetricDifference(l2, l1))
	fmt.Println(SymmetricDifference(l3, l1))
	fmt.Println(SymmetricDifference(l3, l2))

	// Output:
	// [30 40 15 5]
	// [25 -5 35 25 20 0 10]
	// [15 25 5 -5 35 5 25 0 10 20 30 40]
}

func ExampleSymmetricDifference_special_cases() {
	l1 := []int{20, 0, 15, 5, 10}
	l2 := []int{}

	fmt.Println(SymmetricDifference(l1, l2))
	fmt.Println(SymmetricDifference(l2, l1))
	fmt.Println(SymmetricDifference(l1, l1))
	fmt.Println(SymmetricDifference(l2, l2))

	// Output:
	// [20 0 15 5 10]
	// [20 0 15 5 10]
	// []
	// []
}
