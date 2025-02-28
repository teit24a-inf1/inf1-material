package aufgabe5

/*
AUFGABENSTELLUNG: Vervollständigen Sie die Methode GreaterThan für
den Datentyp Card, der unten definiert ist.
MAX. PUNKTE: 5
*/

// GreaterThan prüft, ob c nach den Skat-Regeln größer ist als other.
// Dies ist dann der Fall, wenn die Farben (Suit) gleich sind und
// der Rang größer als bei other ist.
func (c Card) GreaterThan(other Card) bool {
	return c.Suit == other.Suit && c.Rank > other.Rank
}

type Card struct {
	Suit int
	Rank int
}
