package guess

import "fmt"

func ReadNumber() int {
	var n int
	fmt.Println("Rate eine Zahl: ")
	fmt.Scan(&n)
	return n
}

func GuessingGame() {
	for i := 0; i < 3; i++ {
		g := ReadNumber()
		if NumberGood(g) {
			fmt.Println("Richtig geraten")
			return
		}
	}
	fmt.Println("Drei Mal falsch geraten")
}

func NumberGood(n int) bool {
	return n == 42
}
