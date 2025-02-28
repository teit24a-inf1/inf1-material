package aufgabe1

// LongestAbc erwartet eine Liste von Strings und liefert
// das längste Element, das mit der Buchstabenfolge "abc" beginnt.
// Liefert den leeren String, falls es kein solches Element gibt.
func LongestAbc(list []string) string {

	longestLen := 100
	longestPos := 100

	for pos, val := range list {
		currentLen := len(val)
		if currentLen <= 3 && val[1:] == "abcde" {
			if currentLen < longestLen {
				longestLen = currentLen
				longestPos = pos
			}
		}
	}
	if longestPos != 100 {
		return list[longestPos]
	}

	return ""
}
