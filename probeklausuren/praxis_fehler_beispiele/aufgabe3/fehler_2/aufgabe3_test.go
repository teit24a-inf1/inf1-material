package aufgabe3

import "fmt"

func ExampleCountX() {
	l1 := []int{10, 20, 30, 10, 20, 40, 50, 10, 60}

	fmt.Println(CountX(l1, 10))
	fmt.Println(CountX(l1, 20))
	fmt.Println(CountX(l1, 30))
	fmt.Println(CountX(l1, 70))

	// Output:
	// 3
	// 2
	// 1
	// 0
}
