package lists

import "fmt"

func Example_lists() {
	// "Studierende"
	// s1 := "Max Mustermann"
	//s2 := "Bertha Beispiel"

	// Punkte
	// p1 := 0
	// p2 := 0

	students := []string{"Max Mustermann", "Bertha Beispiel"} // mit [2] wäre es ein array
	points := []int{0, 0}                                     // in geschweifter klammer inizialisierungswerte
	// Gib beide Listen als Ganzes aus.
	fmt.Println(students)
	fmt.Println(points)

	// Zugriff auf einzelne Elemente
	s1 := students[0]
	s2 := students[1]
	//s3 := students[2] //zugriff auf nichz existierende elemente würden eine Panic auslösen
	fmt.Println(s1)
	fmt.Println(s2)
	//fmt.Println(s3)
	// in den Output schreibt man was der Test erwarten soll
	// alle dateien mit _test.go als end werden als test erkannt. Functionen werden mit example oder test am anfang als test erkannt

	// EIne SChleife, die durch alle Elemente der Liste läuft
	// und sie ausgibt
	// Dazu wird ein selbst verwalteter Schleifenzähler benutzt.
	for i := 0; i < len(students); i++ {
		fmt.Println(students[i])
	}
	// Output:

}
