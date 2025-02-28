package aufgabe1

// LongestAbc erwartet eine Liste von Strings und liefert
// das längste Element, das mit der Buchstabenfolge "abc" beginnt.
// Liefert den leeren String, falls es kein solches Element gibt.
func LongestAbc(list []string) string {
	return find LongestABC(list, "") // Funktionsaufruf falsch
}
func findLongestABC (list []string, longest string) string {
	if len(list) == 0 { // Fehler verschwindet, wenn oben richtig.
		return Longest
	}
	if len (list[0]) >= 3 && list [0] [3] == "abc" { // Typfehler (wie vom Compiler mitgetgeilt)
		if len (list[0]) > len(longest) {
			longest =list [0]
		}
	}
}
