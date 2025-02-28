package aufgabe6

// import "slices"

// Duplicates erwartet eine int-Liste list.
// Die Funktion liefert eine int-Liste mit den Elementen,
// die mehr als einmal in l1 vorkommen.
// Im Ergebnis kommt jedes Element nur einmal vor.
// Die Elemente stehen im Ergebnis in der Reihenfolge ihres ersten Auftretens.
func Duplicates(list []int) []int {
	seen := make(map[int]bool)
	result := []int{}
	for _, el := range list {
		if !seen[el] {
			seen[el] = true
			result = append(result, el)
		}
	}
	return result
}

//slices.Sort(list)
//for i := 0; i < len(list); i++ {
//	if list [i] < list [i] {
//		result = append(result, list[i])
//	}
// }
