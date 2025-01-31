package functions

import "fmt"

// ForEach erwartet eine Liste aus Zahlen und eine Funktion f.
// Wendet f auf jedes Element der Liste an.
func ForEach(list []int, f func(int)) {
	for _, el := range list {
		f(el)
	}
}

func printNumber(n int) {
	fmt.Println(n)
}

func printTwice(n int) {
	fmt.Println(n)
	fmt.Println(n)
}

func ExampleForEach() {
	l1 := []int{5, 2, 4, 7, 8, 42}

	ForEach(l1, printNumber)
	ForEach(l1, printTwice)

	// Output:
}

func Example_sum() {
	l1 := []int{5, 2, 4, 7, 8, 42}

	sum := 0
	add := func(n int) {
		sum = sum + n
	}
	ForEach(l1, add)

	fmt.Println(sum)

	// Output:
}

func Example_sum_even() {
	l1 := []int{5, 2, 4, 7, 8, 42}

	// Variable für das Ergebnis definieren.
	sum := 0

	// Eine Funktion definieren, die eine Zahl n erwartet
	// und sie genau dann auf sum addiert, wenn n gerade ist.
	add := func(n int) {
		if n%2 == 0 {
			sum = sum + n
		}
	}

	// Die Funktion auf alle Elemente der Liste anwenden.
	ForEach(l1, add)

	// Ergebnis ausgeben.
	fmt.Println(sum)

	// Output:
}

func Example_filter_even() {
	l1 := []int{5, 2, 4, 7, 8, 42}

	// Variable für das Ergebnis definieren.

	// Eine Funktion definieren, die eine Zahl n erwartet
	// und sie genau dann auf sum addiert, wenn n gerade ist.

	// Die Funktion auf alle Elemente der Liste anwenden.

	// Ergebnis ausgeben.

	// Output:
	// [5,7]
}
