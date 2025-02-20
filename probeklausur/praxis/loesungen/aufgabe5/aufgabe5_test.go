package aufgabe5

import "fmt"

func ExampleCard_GreaterThan() {
	c1 := Card{2, 2}
	c2 := Card{2, 3}
	c3 := Card{3, 3}

	fmt.Println(c1.GreaterThan(c2))
	fmt.Println(c2.GreaterThan(c1))
	fmt.Println(c1.GreaterThan(c3))
	fmt.Println(c3.GreaterThan(c1))
	fmt.Println(c2.GreaterThan(c3))
	fmt.Println(c3.GreaterThan(c2))

	// Output:
	// false
	// true
	// false
	// false
	// false
	// false
}
