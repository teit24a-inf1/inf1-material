package structs

import "fmt"

func ExamplePerson_zugriff() {
	p1 := NewPerson("Bertha Beispiel", 23)

	fmt.Println(p1)

	p1.Alter = 77
	fmt.Println(p1)

	p1.Name = "Arno Nym"
	fmt.Println(p1)

	p1.Phone = PhoneNumber{
		AreaCode:    "176",
		CountryCode: "+49",
		Number:      "1234567",
	}
	fmt.Println(p1)
	fmt.Println(p1.Phone.AreaCode)

	// Output:
}
