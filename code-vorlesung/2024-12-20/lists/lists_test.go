package lists

import "fmt"

func Example_lists() {
	// "Studierende"
	// s1 := "Max Mustermann"
	//s2 := "Bertha Beispiel"

	// Punkte
	// p1 := 0
	// p2 := 0

	students := []string{"Max Mustermann", "Bertha Beispiel"}
	points := []int{0, 0}
	// in geschweifter klammer inizialisierungswerte
	fmt.Println(students)
	fmt.Println(points)
	// in den Output schreibt man was der Test erwarten soll
	// alle dateien mit _test.go als end werden als test erkannt. Functionen werden mit example oder test am anfang als test erkannt
	// Output:
	//[Max Mustermann Bertha Beispiel]
	//[0 0]
}
