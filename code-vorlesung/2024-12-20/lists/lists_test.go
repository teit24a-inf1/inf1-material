package lists

import "fmt"

func Example_lists() {
<<<<<<< HEAD
	// "Studierende"
	// s1 := "Max Mustermann"
	//s2 := "Bertha Beispiel"
=======
	// Studierende
	// s1 := "Max Mustermann"
	// s2 := "Bertha Beispiel"
>>>>>>> 7d8a5d8c0ae333e4b23ae83f6bee7832b1ce11ac

	// Punkte
	// p1 := 0
	// p2 := 0

<<<<<<< HEAD
	students := []string{"Max Mustermann", "Bertha Beispiel"} // mit [2] wäre es ein array
	points := []int{0, 0}                                     // in geschweifter klammer inizialisierungswerte
=======
	// Definiere zwei Listen.
	students := []string{"Max Mustermann", "Bertha Beispiel"}
	points := []int{0, 0, 13, 25, 0, 2, -15, 7}

>>>>>>> 7d8a5d8c0ae333e4b23ae83f6bee7832b1ce11ac
	// Gib beide Listen als Ganzes aus.
	fmt.Println(students)
	fmt.Println(points)

	// Zugriff auf einzelne Elemente
	s1 := students[0]
	s2 := students[1]
<<<<<<< HEAD
	//s3 := students[2] //zugriff auf nichz existierende elemente würden eine Panic auslösen
	fmt.Println(s1)
	fmt.Println(s2)
	//fmt.Println(s3)
	// in den Output schreibt man was der Test erwarten soll
	// alle dateien mit _test.go als end werden als test erkannt. Functionen werden mit example oder test am anfang als test erkannt

	// EIne SChleife, die durch alle Elemente der Liste läuft
	// und sie ausgibt
=======
	fmt.Println(s1)
	fmt.Println(s2)

	// Zugriff auf nicht existierende Elemente
	// würde eine Panic auslösen.
	//s3 := students[2]
	//fmt.Println(s3)

	// Eine Schleife, die durch alle Elemente der Liste students
	// läuft und sie ausgibt.
>>>>>>> 7d8a5d8c0ae333e4b23ae83f6bee7832b1ce11ac
	// Dazu wird ein selbst verwalteter Schleifenzähler benutzt.
	for i := 0; i < len(students); i++ {
		fmt.Println(students[i])
	}
<<<<<<< HEAD
=======

	// Eine Schleife, die durch alle Elemente der Liste points
	// läuft und sie ausgibt.
	// Dazu wird ein selbst verwalteter Schleifenzähler benutzt.
	for i := 0; i < len(points); i = i + 2 {
		fmt.Println(points[i])
	}

	// Eine Schleife, die prüft, ob points aufsteigend sortiert ist.
	for i := 0; i < len(points)-1; i++ {
		if points[i] > points[i+1] {
			fmt.Println("Points ist nicht sortiert:")
			fmt.Printf("points[%v] > points[%v]\n", i, i+1)
			fmt.Printf("points[%v] == %v\n", i, points[i])
			fmt.Printf("points[%v] == %v\n", i+1, points[i+1])
			break
		}
	}

	// Eine Schleife, die die Liste points ohne selbst
	// verwalteten Schleifenzähler durchläuft.
	for i, el := range points {
		fmt.Printf("%v: %v\n", i, el)
	}

	// Eine Schleife, die die Liste points bis zur Stelle 4 läuft.
	for i, el := range points[:5] {
		fmt.Printf("%v: %v\n", i, el)
	}

>>>>>>> 7d8a5d8c0ae333e4b23ae83f6bee7832b1ce11ac
	// Output:

}
