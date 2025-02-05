package main

import "fmt"

func main() {
	c1 := NewColourCode(101, "Red")
	c1.anotherString = "Full Red"
	c1.Display()
}

func NewColourCode(code int, name string) *ColourCode {
	return &ColourCode{int: code, string: name}
}

type anotherString = string // This is not a new type, it is just an alias

type ColourCode struct {
	int           // anonymous field or embedded field
	string        // anonymous field or embedded field
	anotherString // anotherString is not a new filed , it is just an alias for string.Since the name is different ,it can be used as a field
}

func (cc *ColourCode) Display() {
	fmt.Println("Code:", cc.int)
	fmt.Println("Colour:", cc.string)
	fmt.Println("Full Colour:", cc.anotherString)
}
