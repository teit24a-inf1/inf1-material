package digits

import "fmt"

func ExampleDigits() {
	fmt.Println(Digits(42))
	fmt.Println(Digits(123))

	// Output:
	// [4 2]
	// [1 2 3]
}
