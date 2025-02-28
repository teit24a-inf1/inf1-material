package aufgabe1

import (
	"fmt"
)

func ExampleLongestAbc() {
	l1 := []string{"abcd", "ab", "de", "abcde", "defabcedfg", "kj"}
	l2 := []string{"abcd", "ab", "de", "abc", "defabcedfg", "kj"}
	l3 := []string{"dabc", "ab", "de", "cabcde", "defabcedfg", "kj"}
	l4 := []string{"abc"}

	fmt.Println(LongestAbc(l1))
	fmt.Println(LongestAbc(l2))
	fmt.Println(LongestAbc(l3))
	fmt.Println(LongestAbc(l4))

	// Output:
	// abcde
	// abcd
	//
	// abc
}
