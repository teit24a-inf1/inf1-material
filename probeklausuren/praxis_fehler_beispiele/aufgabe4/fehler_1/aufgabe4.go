package aufgabe4

// ElementProducts erwartet zwei int-Listen l1 und l2.
// Sie liefert eine Liste, die an jeder Position
// jeweils das Produkt der beiden Elemente enthält.
// Falls eine Position nur in einer Liste vorkommt,
// soll dieses Element ins Ergebnis übernommen werden.
func ElementProducts(l1, l2 []int) []int {
	result := []int{}

	if len(l1) == 0 {
		return l2
	}

	if len(l2) == 0 {
		return l1
	}

	for _, el1 := range l1 {
		if Add(l2, el1) {
			result = append(result, Add(addRes))
		}
	}

	for _, el2 := range l2 {
		if Add(l1, el2) {
			result = append(result, Add(addRes))
		}
	}

	return []int{}
}

func Add(l []int, el int) {
	for _, e := range l {
		addRes := e * el
		return addRes
	}
	return 0
}
