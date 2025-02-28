package aufgabe2

import "fmt"

func ExampleExcludeBetween() {
	l1 := []int{10, 30, 40, 60, 70, 80, 90, 50, 20}

	fmt.Println(ExcludeBetween(l1, 20, 50))
	fmt.Println(ExcludeBetween(l1, 19, 51))
	fmt.Println(ExcludeBetween(l1, 50, 20))

	// Output:
	// [30 40]
	// [20 30 40 50]
	// []
}
