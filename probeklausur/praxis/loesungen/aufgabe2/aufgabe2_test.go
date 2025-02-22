package aufgabe2

import "fmt"

func ExampleExcludeStringsBetween() {
	l1 := []string{"abc", "BEGIN", "def", "END", "ghi"}
	l2 := []string{"BEGIN", "def", "END", "ghi"}
	l3 := []string{"BEGIN", "def", "ghi", "END"}
	l4 := []string{"END", "def", "BEGIN", "ghi"}
	l5 := []string{"abc", "def", "ghi"}

	fmt.Println(ExcludeStringsBetween(l1, "BEGIN", "END"))
	fmt.Println(ExcludeStringsBetween(l2, "BEGIN", "END"))
	fmt.Println(ExcludeStringsBetween(l3, "BEGIN", "END"))
	fmt.Println(ExcludeStringsBetween(l4, "BEGIN", "END"))
	fmt.Println(ExcludeStringsBetween(l5, "BEGIN", "END"))

	// Output:
	// [abc ghi]
	// [ghi]
	// []
	// []
	// []
}
