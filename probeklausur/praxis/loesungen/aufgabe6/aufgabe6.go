package aufgabe6

/*
AUFGABENSTELLUNG: Vervollständigen Sie die Funktion SymmetricDifference.
MAX. PUNKTE: 10
*/

// SymmetricDifference erwartet zwei int-Listen l1 und l2.
// Die Funktion liefert eine int-Liste mit den Elementen,
// die in einer, aber nicht in beiden Listen vorhanden sind.
//
// Die Elemente kommen dabei in der gleichen Reihenfolge vor, wie in den
// Ursprungslisten, mehrfach vorkommende Elemente werden entsprechend wiederholt.
// Die Elemente aus l1 kommen vor denen aus l2 in der Ergebnisliste vor.
func SymmetricDifference(l1, l2 []int) []int {
	result := []int{}
	for _, el1 := range l1 {
		if !Contains(l2, el1) {
			result = append(result, el1)
		}
	}
	for _, el2 := range l2 {
		if !Contains(l1, el2) {
			result = append(result, el2)
		}
	}
	return result
}

func Contains(l []int, el int) bool {
	for _, e := range l {
		if e == el {
			return true
		}
	}
	return false
}
