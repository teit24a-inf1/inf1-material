package aufgabe2

/* AUFGABENSTELLUNG: Vervollständigen Sie die Funktion FilterVowels.
 * ERREICHBARE PUNKTE: 10
 */

// FilterVowels erwartet einen String s und liefert einen String zurück,
// der aus s entsteht, indem alle Vokale entfernt werden.
// Anmerkung: Es müssen nur Kleinbuchstaben beachtet werden.
func FilterVowels(s string) string {
	result := ""
	vowels := "aeiou"
	for _, char := range s{
		if !contains(vowels, char){
			result += string(char)
		}
	}
	return result


c contains (set string, char rune) bool{
	for _, c := range set {
		if c == char {
			return true
		}
	}
	return false 
}
