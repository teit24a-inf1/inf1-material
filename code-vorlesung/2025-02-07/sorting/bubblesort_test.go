package sorting

import "fmt"

func ExampleBubbleSort() {
	l1 := []int{17, 2, 3, 25, 32, 38}

	BubbleSort(l1)

	fmt.Println(l1)

	// Output:
	// [2 3 17 25 32 38]
}
