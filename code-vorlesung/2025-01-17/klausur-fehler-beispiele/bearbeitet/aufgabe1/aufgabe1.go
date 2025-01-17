package aufgabe1

/* AUFGABENSTELLUNG: Vervollständigen Sie die Funktion EvenPrefix.
 * ERREICHBARE PUNKTE: 10
 */

// EvenPrefix erwartet eine Liste von Zahlen und liefert
// die längste Anfangs-Folge, die nur gerade Zahlen enthält.
// D.h. eine Liste mit allen geraden Zahlen vor der ersten ungeraden.
func EvenPrefix(list []int) []int {
	result := []int{}
	for_, num := range list {
		if num%2 == 0 {
			result = append(result, num)
		} else {
			break
		}
	}
	return result
}
