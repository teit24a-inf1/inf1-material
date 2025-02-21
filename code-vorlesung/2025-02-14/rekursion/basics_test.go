package main

import "fmt"

func Example_count() {

	CountDown(3)
	CountUp(3)

	// Output:
}

func ExampleSubtract() {

	fmt.Println(Subtract(3))

	// Output:
}

func ExamplePower5() {
	fmt.Println(Power5(-5))
	fmt.Println(Power5(0))
	fmt.Println(Power5(1))
	fmt.Println(Power5(2))
	fmt.Println(Power5(3))
	fmt.Println(Power5(4))
	fmt.Println(Power5(5))

	// Output:
}

// Zusatz-Überlegung

// func Foo(n int) {
// 	fmt.Println(n)
// 	Bar(n - 1)
// 	fmt.Println(n)
// }

// func Bar(x int) {
// 	x = 2 * x
// 	fmt.Println(x)
// }
