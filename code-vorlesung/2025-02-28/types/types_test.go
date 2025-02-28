package types

import "fmt"

func Example_int_types() {

	x1 := 42
	l1 := []int{1, 2, 3, 4, 5}

	fmt.Printf("x1: %T\n", x1)
	fmt.Printf("l1: %T\n", l1)
	fmt.Printf("l1[0]: %T\n", l1[0])

	// Output:
}

func Foo() string {
	return "abc"
}

func Example_var_declarations() {
	x1 := 42.0
	x2 := 42
	x3 := Foo()

	fmt.Printf("x1: %T\n", x1)
	fmt.Printf("x2: %T\n", x2)
	fmt.Printf("x3: %T\n", x3)

	// Output:
}

func Example_loop_counters() {
	l1 := []int{5, 4, 3, 2, 1}
	l2 := []string{"5", "4", "3", "2", "1"}

	for i, v := range l1 {
		fmt.Printf("%d: %d\n", i, v)
	}
	fmt.Println()
	for i, v := range l2 {
		fmt.Printf("%d: %s\n", i, v)
	}
	fmt.Println()
	for i, v := range "abcde" {
		fmt.Printf("%d: %v\n", i, v) // Falsche Ausgabe: Unicode-Werte statt Buchstaben
		// fmt.Printf("%d: %c\n", i, v) // Korrektur
	}

	// Output:

}

func Example_loop_slice() {
	l1 := []int{5, 4, 3, 2, 1}

	for i, v := range l1[1:] {
		fmt.Printf("%d: %d (%d)\n", i, v, l1[i])
	}

	// Output:
}
