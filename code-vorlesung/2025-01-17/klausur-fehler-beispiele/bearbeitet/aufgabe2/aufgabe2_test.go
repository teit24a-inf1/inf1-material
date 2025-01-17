package aufgabe2

import "fmt"

func ExampleFilterVowels() {

	fmt.Println(FilterVowels("abcde"))
	fmt.Println(FilterVowels("aeiou"))
	fmt.Println(FilterVowels("xyz"))

	// Output:
	// bcd
	//
	// xyz
}
