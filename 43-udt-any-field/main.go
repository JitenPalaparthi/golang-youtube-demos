package main

import "fmt"

func main() {

	p1 := NewPerson("Jiten", "Jiten@Gmail.Com", "active", "Bangalore-1,India, 560086")
	// details are just a string
	fmt.Println(p1)

	p2 := NewPerson("Priya", "Priya@Gmail.Com", "active", struct{ Line1, City, PinCode string }{Line1: "RajajiNagar-1", City: "Bangalore", PinCode: "560086"})
	fmt.Println(p2)
	p3 := NewPerson("Raja", "Raja@Gmail.Com", "active", struct{ Twitter, Linkedin string }{Twitter: "raja_1", Linkedin: "linkedin.com/in/raja101"})
	fmt.Println(p3)
}

func NewPerson(name, email, status string, details any) *Person {
	return &Person{Name: name, Email: email, Status: status, Details: details}
}

type Person struct {
	Name, Email, Status string
	Details             any // can assign any kind of value
}

// Create a display function
// If any is a type of a struct{Line1,City,PinCode string} print that
// if any is a type of struct{Twitter,Linkedin string} print that
// if any is not of those two types .. print nothing w.r.t details
