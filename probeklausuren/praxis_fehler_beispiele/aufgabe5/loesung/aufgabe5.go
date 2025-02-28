package aufgabe5

// IsChain erwartet eine Liste von Dominoe-Objekten.
// Die Funktion prüft, ob diese Liste eine Kette bildet,
// die nach den Domino-Regeln erlaubt wäre.
// Bei einer solchen Kette muss immer die rechte Seite eines Steins
// gleich der linken Seite des nächsten Steins sein.
func IsChain(dominoes []Dominoe) bool {
	for i := 0; i+1 < len(dominoes); i++ {

		if dominoes[i].Right != dominoes[i+1].Left {
			return false
		}
	}
	return true
}

// Dominoe repräsentiert einen Domino-Stein mit zwei Zahlen.
type Dominoe struct {
	Left  int
	Right int
}
