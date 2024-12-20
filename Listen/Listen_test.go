package lists

import "fmt"

func Example_lists() {

	//Studierende
	//s1 := "Max Mustermann"
	//s2 := "Bertha Beispiel"

	//Punkte
	//p1 := 0
	//p2 := 0

	//[] definiert die Funktion als Liste (Leeres klammernpaar)
	//{} initialisierung der Liste (Leere Liste)

	//definierung zweier Listen
	students := []string{"Max Musterman", "Bertha Beispiel"}
	points := []int{0, 0}

	fmt.Println(students)
	fmt.Println(points)

	//Zugriff auf einzelne Elemente
	s1 := students[0]
	s2 := students[1]
	//s3 := students[2]
	fmt.Println(s1)
	fmt.Println(s2)
	//fmt.Println(s3)

	//Iene Schleife,die durch alle Elemente der Liste students
	//läuft und sie ausgibt.
	//Dazu wird eine selbst verwalter Schleifenzähler benutzt.
	for i := 0; i < len(students); i++ {
		fmt.Println(students[i])
	}

	//Iene Schleife,die durch alle Elemente der Liste points
	//läuft und sie ausgibt.
	//Dazu wird eine selbst verwalter Schleifenzähler benutzt.
	for i := 0; i < len(points); i = i + 2 {
		fmt.Println(points[i])
	}

	//Eine schleife, die Prüft, ob Points aufsteigend sortiert ist.
	for i := 0; i < len(points); i++ {
		if points[i] > points[i+1] {
			fmt.Printf("Points ist nicht sortiert:")
			fmt.Printf("Points [%v] >  points[%v]\n", i, i+1)
			fmt.Printf("Points [%v] == %v\r", i, points[i])
			fmt.Printf("points[%v] == %v\r", i+1, points[i+1])
			break
		}
	}

	//Eine Schleife, die die Liste points ohne selbst verwalten Schleifenzähler durchläuft.
	for i, el := range points {
		fmt.Printf("%v: %v\n", i, el)
	}

	//Eine Schleife, die die Liste points bis zur stelle 4 läuft
	for i, el := range points[0:5] {
		fmt.Printf("%v: %v\n", i, el)
	}

	//Output:

}
