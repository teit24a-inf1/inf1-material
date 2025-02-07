package sorting

func QuickSort(list []int) {
	// Listen der Länge 0 oder 1 sind schon sortiert.
	if len(list) <= 1 {
		return
	}

	// Pivot-Element wählen:
	p := list[0]

	// Liste durchlaufen, kleinere Elemente als p in die linke Teilliste,
	// größere in die rechte Teilliste.
	links := []int{}
	rechts := []int{}

	for _, el := range list[1:] {
		if el < p {
			links = append(links, el)
		} else {
			rechts = append(rechts, el)
		}
	}

	// links und rechts sortieren
	QuickSort(links)
	QuickSort(rechts)

	// Die listen wieder zusammenkopieren.
	result := []int{}
	result = append(result, links...)
	result = append(result, p)
	result = append(result, rechts...)
	copy(list, result)

}
