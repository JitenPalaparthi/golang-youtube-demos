package main

import "fmt"

func main() {

	s1 := NewStudent(101, "Jiten", "Jiten@Gmail.Com", "active", &Address{Line1: "Rajaji Nagar Street1", State: "Karnataka", City: "Bengaluru", Country: "India", PinCode: "560086"}, &SocialMedia{Twitter: "jiten_1", Linkedin: "linkedin.in/jpalaparthi"})

	fmt.Println("City:", s1.City)

	s1.Display()
	s1.Address.Display() // Since there is same name used for both , then call them specifically with the type

}

func NewStudent(no int, name, email, status string, address *Address, socialMedia *SocialMedia) *Student {
	return &Student{No: no, Name: name, Email: email, Address: address, SocialMedia: socialMedia}
}

type Student struct {
	No           int
	Name, Email  string
	Status       string
	*Address     // Promoted field , can access the members (fields and methods) of the promoted filed type directly
	*SocialMedia // Promoted filed
}

type SocialMedia struct {
	Twitter  string
	Linkedin string
}

func (s *Student) Display() {
	fmt.Println("Student")
	fmt.Println("No:", s.No)
	fmt.Println("Name:", s.Name)
	fmt.Println("Email:", s.Email)
	fmt.Println("Status:", s.Status)

	if s.SocialMedia != nil {
		fmt.Println("Social Media")
		fmt.Println("Twitter:", s.Twitter)
		fmt.Println("Linkedin:", s.Linkedin)
	}
}

type Address struct {
	Line1   string
	City    string
	State   string
	Country string
	PinCode string
}

func (a *Address) Display() {
	fmt.Println("Address")
	fmt.Println("Line1:", a.Line1)
	fmt.Println("City:", a.City)
	fmt.Println("State:", a.State)
	fmt.Println("Country:", a.Country)
	fmt.Println("PinCode:", a.PinCode)
}
