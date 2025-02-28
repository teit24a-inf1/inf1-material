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
	firstpos := -1
	lastpos := -1
	for pos, s := range list {
		if s == first {
			firstpos = pos
		}
		if s == last {
			lastpos = pos
		}
	}
	if lastpos <= firstpos {
		return []string{}
	}
	return append(list[:firstpos], list[lastpos+1:]...)
}
