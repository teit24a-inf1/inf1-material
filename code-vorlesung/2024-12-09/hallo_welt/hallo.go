package main

import "fmt"

func main() {
	Greet("Welt")
	Greet("Kurs")
}

func Greet(n string) {
	fmt.Println("Hallo " + n)
}
