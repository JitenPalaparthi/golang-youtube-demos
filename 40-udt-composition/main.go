package main

import "fmt"

func main() {

	a1 := NewAddress("Electronic Street", "Dubai", "Dubai", "UAE", "0001")
	e1 := NewEmployee(101, "Jiten", "Jiten@Gmail.Com", "012312213", a1)
	e1.Display()

	e2 := Employee{Id: 102, Name: "Jiten", Email: "Jiten@Gmail.Com", Contact: "101010", Address: &Address{Line1: "Electronic Street", City: "Dubai", Country: "UAE", PinCode: "971"}}

	println()
	e2.Display()
	//e2.Address.City
}

type Employee struct {
	Id                   int
	Name, Email, Contact string
	Address              *Address // The name of the field can be the name of the udt
}

func NewEmployee(id int, name, email, contact string, address *Address) *Employee {
	return &Employee{Id: id, Name: name, Email: email, Contact: contact, Address: address}
}

func (e *Employee) Display() {
	fmt.Println("Id:", e.Id)
	fmt.Println("Name:", e.Name)
	fmt.Println("Email:", e.Email)
	fmt.Println("Contact:", e.Contact)
	if e.Address != nil {
		e.Address.Display()
	}
}

func NewAddress(line1, city, state, country, pinCode string) *Address {
	return &Address{Line1: line1, City: city, State: state, Country: country, PinCode: pinCode}
}

type Address struct {
	Line1, City, State, Country, PinCode string
}

func (a *Address) Display() {
	fmt.Println("Address")
	fmt.Println("Line1:", a.Line1)
	fmt.Println("City:", a.City)
	fmt.Println("State:", a.State)
	fmt.Println("Country:", a.Country)
	fmt.Println("PinCode:", a.PinCode)
}
