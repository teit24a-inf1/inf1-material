package aufgabe3

// CountX erwartet eine Liste von Zahlen sowie eine Zahl x.
// CountX liefert die Anzahl der Vorkommen von x in der Liste.
func CountX(list []int, x int) int {

	counter := 0

	if len(list) == 0 {
		return 0
	}

	if list[0] == x {
		counter++
		return CountX(list[1:], x)
	}

	if list[0] != x {
		return CountX(list[1:], x)
	}

	return counter
}
