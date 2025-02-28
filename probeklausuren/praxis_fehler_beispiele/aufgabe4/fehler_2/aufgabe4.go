package aufgabe4

// ElementProducts erwartet zwei int-Listen l1 und l2.
// Sie liefert eine Liste, die an jeder Position
// jeweils das Produkt der beiden Elemente enthält.
// Falls eine Position nur in einer Liste vorkommt,
// soll dieses Element ins Ergebnis übernommen werden.
func ElementProducts(l1, l2 []int) []int {
	l3 := []int{}
	längel1 := len(l1)
	länge2 := len(l2)
	var vorrübergehend int
	if len(l1) == 0 || len(l2) == 0 {
		return []int{}
	}
	if l1[0] != 0 {
		vorrübergehend = (l1[0] * l2[0])
		l3 = append(l3, vorrübergehend)
	}
	return ElementProducts(l1[1:], l2[1:]) + []int{vorrübergehend}
}
