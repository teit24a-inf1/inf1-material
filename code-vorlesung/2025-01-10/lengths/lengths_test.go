package lengths

import "fmt"

type Length int

func (l Length) Centimeters() int {
	return int(l / 10)
}

func (l Length) Meters() int {
	return l.Centimeters() / 100
}

func (l Length) Kilometers() int {
	return l.Meters() / 1000
}

func FromCentimeters(c int) Length {
	return Length(c * 10)
}

func ExampleLength() {
	//var a Length = 123456789
	a := FromCentimeters(123456789)

	fmt.Println(a)
	fmt.Printf("%T\n", a)

	fmt.Println(a.Centimeters())
	fmt.Println(a.Meters())
	fmt.Println(a.Kilometers())

	// Output:
	// 1234567890
	// lengths.Length
	// 123456789
	// 1234567
	// 1234
}
