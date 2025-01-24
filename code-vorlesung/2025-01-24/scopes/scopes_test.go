package scopes

import "fmt"

func Example_scope() {
	x := Foo()

	fmt.Println(x)

	// Output:
}

func Foo() int {
	x := 42
	Bar(&x)
	y := 23
	z := x + y

	return z
}

func Bar(x *int) {
	*x = *x + 2
}

// Globale Variable zu Beispielzwecken.
// Bitte auf keinen Fall in Produktivcode benutzen!
var x int = 1

func Example_scope_2() {
	x *= 2
	fmt.Println(x)
	x := 10
	{
		x := 20
		y := 42
		{
			x := 30
			fmt.Println(x)
			fmt.Println(y)
		}
		fmt.Println(x)
	}
	fmt.Println(x)

	// Output:
}
