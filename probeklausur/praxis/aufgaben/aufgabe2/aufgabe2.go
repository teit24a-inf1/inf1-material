package aufgabe2

/*
AUFGABENSTELLUNG: Vervollständigen Sie die Funktion ExcludeStringsBetween.
MAX. PUNKTE: 10
*/

// ExcludeStringsBetween erwartet eine Liste und zwei Strings first und last.
// Die Funktion liefert eine Liste mit allen Elementen, die nicht zwischen first und last liegen.
// first und last sollen nicht zum Ergebnis gehören.
// Falls die Liste first oder last nicht enthält, oder falls last vor first vorkommt,
// soll die leere Liste geliefert werden.
func ExcludeStringsBetween(list []string, first, last string) []string {
	firstint := -1
	lastint := -1
	for i := 0; i < len(list); i++ {
		if list[i] == first {
			firstint = i
			break
		}
	}
	for t := 0; t < len(list); t++ {
		if list[t] == last {
			lastint = t
			break
		}
	}
	if firstint == -1 || lastint == -1 {
		return []string{}
	}
	if firstint >= lastint {
		return []string{}
	}
	return append(list[:firstint], list[lastint+1:]...)
}
