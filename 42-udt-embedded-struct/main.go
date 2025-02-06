package main

import "fmt"

func main() {

	p1 := NewPerson("Jiten", "Jiten@Gmail.Com", "active", struct {
		Line1   string
		City    string
		State   string
		Country string
		PinCode string
	}{Line1: "abcd1", City: "Dubai", State: "Dubai", Country: "UAE", PinCode: "971"}, struct {
		Twitter  string
		Linkedin string
	}{Twitter: "jitenp", Linkedin: "linkedin.com/in/jpalaparthi"})

	p1.Display()

	p2 := &Person{
		Name:   "Jiten",
		Email:  "Jiten@Gmail.com",
		Status: "active",
		Address: struct {
			Line1   string
			City    string
			State   string
			Country string
			PinCode string
		}{Line1: "abcd1", City: "Dubai", State: "Dubai", Country: "UAE", PinCode: "971"},
		SocialMedia: struct {
			Twitter  string
			Linkedin string
		}{Twitter: "jitenp", Linkedin: "linkedin.com/in/jpalaparthi"},
	}

	p2.Name = "Priya"
	p2.Display()

	p3 := new(Person) // using new builtin function

	p3.Name = "Rahim"

	p3.Display()

}

type Person struct {
	Name, Email string
	Status      string
	Address     struct {
		Line1, City, State, Country, PinCode string
	} // embedded struct
	SocialMedia struct {
		Twitter  string
		Linkedin string
	}
}

func NewPerson(name, email, status string, address struct {
	Line1, City, State, Country, PinCode string
}, socialMedia struct {
	Twitter  string
	Linkedin string
}) *Person {
	return &Person{
		Name: name, Email: email, Status: status, Address: address, SocialMedia: socialMedia,
	}

}
func (p *Person) Display() {
	fmt.Println("Person")
	fmt.Println("Name:", p.Name)
	fmt.Println("Email:", p.Email)
	fmt.Println("Status:", p.Status)

	fmt.Println("Address")
	fmt.Println("Line1:", p.Address.Line1)
	fmt.Println("City:", p.Address.City)
	fmt.Println("State:", p.Address.State)
	fmt.Println("Country:", p.Address.Country)
	fmt.Println("PinCode:", p.Address.PinCode)

	fmt.Println("Twitter:", p.SocialMedia.Twitter)
	fmt.Println("Linkedin:", p.SocialMedia.Linkedin)

}
