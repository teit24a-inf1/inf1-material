package aufgabe5

import "fmt"

func ExampleIsChain() {
	d1 := []Dominoe{{1, 2}, {2, 3}, {3, 4}}
	d2 := []Dominoe{{1, 2}, {2, 3}, {4, 5}}
	d3 := []Dominoe{{1, 2}, {2, 3}, {3, 5}}
	d4 := []Dominoe{{5, 4}, {4, 3}, {3, 2}, {2, 1}}
	d5 := []Dominoe{}

	fmt.Println(IsChain(d1))
	fmt.Println(IsChain(d2))
	fmt.Println(IsChain(d3))
	fmt.Println(IsChain(d4))
	fmt.Println(IsChain(d5))

	// Output:
	// true
	// false
	// true
	// true
	// true
}
