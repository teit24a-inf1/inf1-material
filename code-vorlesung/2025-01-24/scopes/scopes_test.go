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

func Example_scope_3() {
	i := 0

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	fmt.Println(i)

	// Output:
}

func Example_scope_4() {
	i := 0

	for i < 10 {
		fmt.Println(i)
		i++
	}
	fmt.Println(i)

	// Output:
}
