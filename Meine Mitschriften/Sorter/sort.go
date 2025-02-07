package Sorter

func BubbleSort(list []int) {
	for j := 0; j < len(list); j++ {
		//für jede position in der Liste:
		for i := 0; i+1 < len(list); i++ {
			//wenn das Element an stelle i größer isl
			// als das an i+1 (also der rechte Nacbar)...
			if list[i] > list[i+1] {
				//Vertausche die Elemente
				list[i], list[i+1] = list[i+1], list[i]
			}
		}
	}

	//Wenn wir die obige schleife j mal ausgeführt haben, sind die j größten Elemente sortiert am Ende der Liste

}

//die archaische Methode zum Vertauschen von Elementen
//if list[0] > list[1] {
//	h := list[0]
//	list[0] = list[1]
//	list[1] = h
//	}
