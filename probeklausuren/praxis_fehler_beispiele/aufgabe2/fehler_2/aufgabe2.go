package aufgabe2

// ExcludeBetween erwartet eine Liste und zwei Zahlen m und n.
// Die Funktion liefert eine Liste mit allen Elementen x, für die gilt: m < x < n.
func ExcludeBetween(list []int, m, n int) []int {
	result := []int{}
	x := int{}
	for i := len(list); i < x; i++ {
		if m < x < n {
			return x
		}
	}
	return result
}
