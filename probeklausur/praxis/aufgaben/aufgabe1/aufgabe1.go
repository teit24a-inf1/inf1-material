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
	shortestlen := 999999
	shortest := ""
	for i, el := range list {
		if len(el) >= 3 {
			if el[:3] == "abc" {
				if len(el) <= shortestlen {
					shortestlen = len(el)
					shortest = list[i]
				}
			}
		}
	}
	return shortest
}
