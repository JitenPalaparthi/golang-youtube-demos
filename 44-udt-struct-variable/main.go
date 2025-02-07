package main

import "fmt"

func main() {
	var p1 struct{ Name, Email, ContactNo string }
	p1 = struct {
		Name      string
		Email     string
		ContactNo string
	}{Name: "Jiten", Email: "Jiten@Gmail.Com", ContactNo: "9618558500"}
	Display(p1)
}

func Display(p struct{ Name, Email, ContactNo string }) {
	fmt.Println("Person Details")
	fmt.Println("Name:", p.Name)
	fmt.Println("Email:", p.Email)
	fmt.Println("Contact Number:", p.ContactNo)
}
