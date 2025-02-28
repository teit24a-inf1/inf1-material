package function_calls

import "fmt"

func Foo(x int) {
	fmt.Println(x)
}

func Example() {
	Foo(42)

	x := 23
	Foo(x) // richtig
	// Foo(x := 45) // falsch
	// Foo(x int) // falsch

	y := 77
	Foo(y)

	// Output:
}
