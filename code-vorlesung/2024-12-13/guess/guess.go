package guess

import (
	"fmt"
	"math/rand"
)

func ReadNumber() int {
	var n int
	fmt.Println("Rate eine Zahl: ")
	fmt.Scan(&n)
	return n
}

func GuessingGame() {
	// Verwende einen Zufallsgenerator,
	// um die erwartete Zahl zu bestimmen.
	expectedGuess := rand.Intn(16) + 1
	for i := 0; i < 3; i++ {
		g := ReadNumber()
		if NumberGood(expectedGuess, g) {
			fmt.Println("Richtig geraten")
			return
		}
		// Bei falscher Antwort einen Hinweis geben,
		// ob die Zahl zu groß oder zu klein war.
		if expectedGuess < g {
			fmt.Println("Zu groß")
		} else {
			fmt.Println("Zu klein")
		}

	}
	fmt.Println("Drei Mal falsch geraten")
	fmt.Printf("Die richtige Zahl war: %v\n", expectedGuess)
}

func NumberGood(expected, n int) bool {
	return n == expected
}
