package stringbyteslice

import "fmt"

func Example_string() {
	s1 := "abc"
	b1 := []byte(s1)
	s2 := string(b1)

	fmt.Println(s1)
	fmt.Println(b1)
	fmt.Println(s2)

	// Output:
	// abc
	// [97 98 99]
	// abc
}
