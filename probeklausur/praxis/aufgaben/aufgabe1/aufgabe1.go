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
	// TODO
	if len(list)==0{
	return ""}
	länge := 20
	position := 20
	search := "abc"
	for i, el:= range list{
		if len(el)>3 && el[:3]==search{
			länge = len(el)
			position = i
		}
		
	}
}
