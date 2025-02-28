package hanoi

import "fmt"

func Move(s, z string) {
	fmt.Printf("%v -> %v\n", s, z)
}

func Hanoi1(s, m, z string) {
	Move(s, z)
}

func Hanoi2(s, m, z string) {
	Hanoi1(s, z, m)
	Move(s, z)
	Hanoi1(m, s, z)
}
func Hanoi3(s, m, z string) {
	Hanoi2(s, z, m)
	Move(s, z)
	Hanoi2(m, s, z)
}
func Hanoi4(s, m, z string) {
	Hanoi3(s, z, m)
	Move(s, z)
	Hanoi3(m, s, z)
}

func Hanoi(h int, s, m, z string) {
	if h == 0 {
		return
	}

	Hanoi(h-1, s, z, m)
	Move(s, z)
	Hanoi(h-1, m, s, z)

}

// func Hanoi(h int, s, m, z string) {
// 	if h == 1 {
// 		Move(s, z)
// 	} else {
// 		Hanoi(h-1, s, z, m)
// 		Move(s, z)
// 		Hanoi(h-1, m, s, z)
// 	}
// }

func Example_hanoi() {
	Hanoi(2, "A", "B", "C")

	// Output:
}
