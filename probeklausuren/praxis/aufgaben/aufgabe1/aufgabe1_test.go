package aufgabe1

import (
	"fmt"
)

func ExampleShortestAbc() {
	l1 := []string{"abcd", "ab", "de", "abcde", "defabcedfg", "kj"}
	l2 := []string{"abcde", "ab", "de", "abc", "defabcedfg", "kj"}
	l3 := []string{"dabc", "ab", "de", "cabcde", "defabcedfg", "kj"}
	l4 := []string{"abc"}

	fmt.Println(ShortestAbc(l1))
	fmt.Println(ShortestAbc(l2))
	fmt.Println(ShortestAbc(l3))
	fmt.Println(ShortestAbc(l4))

	// Output:
	// abcd
	// abc
	//
	// abc
}
