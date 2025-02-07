package main

import (
	"fmt"
	"reflect"
)

func main() {

	p1 := NewPerson("Jiten", "Jiten@Gmail.Com", "active", "Bangalore-1,India, 560086")
	// details are just a string
	p1.Display()
	fmt.Println("Type of p1.Details:", reflect.TypeOf(p1.Details))
	println()
	p2 := NewPerson("Priya", "Priya@Gmail.Com", "active", struct{ Line1, City, PinCode string }{Line1: "RajajiNagar-1", City: "Bangalore", PinCode: "560086"})
	fmt.Println("Type of p2.Details:", reflect.TypeOf(p2.Details))
	p2.Display()
	println()
	p3 := NewPerson("Raja", "Raja@Gmail.Com", "active", struct{ Twitter, Linkedin string }{Twitter: "raja_1", Linkedin: "linkedin.com/in/raja101"})
	fmt.Println("Type of p3.Details:", reflect.TypeOf(p3.Details))
	p3.Display()

	println()
	p4 := NewPerson("PS", "ps@gmail.com", "active", &Contact{WhatsApp: "9618550011", Mobile: "9618550022", Telephone: "040123456"})
	p4.Display()
	p5 := NewPerson("PS", "ps@gmail.com", "active", struct{ Blog string }{Blog: "www.myblog.com/123"})
	p5.Display() // does not print  struct{ Blog string } since that is not checked in type switch
}

func NewPerson(name, email, status string, details any) *Person {
	return &Person{Name: name, Email: email, Status: status, Details: details}
}

type Person struct {
	Name, Email, Status string
	Details             any // can assign any kind of value
}

type Contact struct {
	WhatsApp, Mobile, Telephone string
}

func (p *Person) Display() {
	fmt.Println("Person")
	fmt.Println("Name:", p.Name)
	fmt.Println("Email:", p.Email)
	fmt.Println("Status:", p.Status)

	switch t1 := p.Details.(type) { // type switch to check the type

	case struct{ Line1, City, PinCode string }:
		fmt.Println("Line1:", t1.Line1)
		fmt.Println("City:", t1.City)
		fmt.Println("PinCode:", t1.PinCode)
	case struct{ Twitter, Linkedin string }:
		fmt.Println("Twitter:", t1.Twitter)
		fmt.Println("Linkedin:", t1.Linkedin)
	case string:
		fmt.Println("Details:", t1)
	case *Contact:
		fmt.Println("WhatsApp:", t1.WhatsApp)
		fmt.Println("Mobile:", t1.Mobile)
		fmt.Println("Telephone:", t1.Telephone)
	}

}

// Create a display function
// If any is a type of a struct{Line1,City,PinCode string} print that
// if any is a type of struct{Twitter,Linkedin string} print that
// if any is not of those two types .. print nothing w.r.t details
