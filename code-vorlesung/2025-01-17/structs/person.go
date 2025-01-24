package structs

import "fmt"

type Person struct {
	Name  string
	Alter int
	Phone PhoneNumber
}

func NewPerson(name string, age int) Person {
	return Person{Name: name, Alter: age, Phone: PhoneNumber{}}
}

type PhoneNumber struct {
	CountryCode string
	AreaCode    string
	Number      string
}

func (p PhoneNumber) String() string {
	return fmt.Sprintf("%s %s %s", p.CountryCode, p.AreaCode, p.Number)
}
