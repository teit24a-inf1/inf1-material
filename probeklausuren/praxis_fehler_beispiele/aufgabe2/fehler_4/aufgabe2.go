package aufgabe2

// ExcludeBetween erwartet eine Liste und zwei Zahlen m und n.
// Die Funktion liefert eine Liste mit allen Elementen x, für die gilt: m < x < n.
func ExcludeBetween(list []int, m, n int) []int {

	mpos := -1
	npos := -1

	for pos, x := range list {
		if x > m {
			mpos = pos
		}
		if x < npos {
			npos = pos
		}
	}

	return append(list[mpos-1:], list[npos])
}
