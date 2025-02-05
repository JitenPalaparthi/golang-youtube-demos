package main

import "fmt"

func main() {
	var p1 Person
	p1.Name = "Jiten"
	p1.Email = "Jiten@Gmail.Com"
	p1.ContactNo = "11100111001"

	fmt.Println("Name:", p1.Name)
	fmt.Println("Email:", p1.Email)
	fmt.Println("Contact Number:", p1.ContactNo)

	var p2 Person = Person{Name: "Jitendra", Email: "Jitendra@Gmail.Com", ContactNo: "122323232"}

	fmt.Println(p2)

	// shorthand declaration
	p3 := Person{Name: "Jiten", Email: "Jiten@Gmail.Com", ContactNo: "10101010"}
	fmt.Println(p3)

	p4 := New("Jiten", "Jiten@Gmail.com", "131231121")
	fmt.Println(p4)
}

type Person struct {
	// Name      string
	// Email     string
	// ContactNo string
	Name, Email, ContactNo string
}

func New(name, email, contactno string) Person {
	return Person{Name: name, Email: email, ContactNo: contactno}
}

func NewPerson(name, email, contactno string) Person {
	return Person{Name: name, Email: email, ContactNo: contactno}
}
