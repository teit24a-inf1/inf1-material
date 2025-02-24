package sorting

func BubbleSort(list []int) {

	for j := 0; j < len(list); j++ {
		// Für jede Position in der Liste:
		for i := 0; i+1 < len(list); i++ {
			// Wenn das Element an Stelle i größer
			// ist als das an i+1 (also als der rechte Nachbar) ...
			if list[i] > list[i+1] {
				// ... vertausche die Elemente
				list[i], list[i+1] = list[i+1], list[i]
			}
		}

		// Wenn wir die obige Schleife j mal ausgeführt haben,
		// sind die j größten Elemente sortiert am Ende der Liste.
		// Also führen wir die Schleife insgesamt len(list) mal aus.
	}

}

// Allgemeines:

// Die archaische Methode zum Vertauschen von Elementen
// 	h := list[0]
// 	list[0] = list[1]
// 	list[1] = h
