package aufgabe1

/*
AUFGABENSTELLUNG: Vervollständigen Sie die Funktion ShortestAbc.
MAX. PUNKTE: 10
*/

// ShortestAbc erwartet eine Liste von Strings und liefert
// das kürzeste Element, das mit der Buchstabenfolge "abc" beginnt.
// Liefert den leeren String, falls es kein solches Element gibt.
//
// Hinweis: Die Funktion muss nur mit kurzen Strings der Länge < 100 funktionieren.
func ShortestAbc(list []string) string {
	shortestLen := 100
	shortestPos := 100
	for pos, val := range list {
		currentLen := len(val)
		if currentLen >= 3 && val[:3] == "abc" {
			if currentLen < shortestLen {
				shortestLen = currentLen
				shortestPos = pos
			}
		}
	}
	if shortestPos != 100 {
		return list[shortestPos]
	}
	return ""
}
