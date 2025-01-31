package functions

import "fmt"

func PrintFoo() {
	fmt.Println("Foo")
}

func PrintBar() {
	fmt.Println("Bar")
}

// CallFunction erwartet eine Funktion als Parameter und ruft sie auf.
// Die Funktion, die erwartet wird, hat selbst keine Parameter und
// liefert nichts zurück.
func CallFunction(f func()) {
	f()
}

func ExamplePrintFoo() {

	PrintBaz := func() {
		fmt.Println("Baz")
	}

	CallFunction(PrintFoo)
	CallFunction(PrintBar)
	CallFunction(PrintBaz)

	printABC := CreatePrinter("ABC")
	printDEF := CreatePrinter("DEF")

	CallFunction(printABC)
	CallFunction(printDEF)

	// Output:
	// Foo
	// Bar
	// Baz
}

func CreatePrinter(s string) func() {
	return func() {
		fmt.Println(s)
	}
}
