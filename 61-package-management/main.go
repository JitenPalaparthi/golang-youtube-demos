package main

import (
	"shapes/shape"
	rectangle "shapes/shape/rect"
	"shapes/shape/square"

	. "math/rand"
	. "time"
)

func main() {
	r1 := rectangle.NewRect(12.3, 13.4)
	// a1 := r1.Area()
	// p1 := r1.Perimeter()
	// fmt.Println("Area of Rect a1:", a1)
	// fmt.Println("Perimeter of Rect p1:", p1)
	//r1.display() // cannot be called. It is restricted
	// r1.Display()
	//r1.a // cannot be accessed
	//r1.p // cannot be accessed
	//rect.Public = "public variable"

	s1 := square.Square(100.12)

	// a2 := s1.Area()
	// p2 := s1.Perimeter()
	// fmt.Println("Area of Square a2:", a2)
	// fmt.Println("Perimeter of Square p2:", p2)
	shape.Shape(r1)
	shape.Shape(s1)

	for i := 1; i <= 10; i++ {
		if n := Intn(100); n%2 == 0 {
			println(n, " is even number. Time is ", Now().Unix())
		}
	}

}

// 1. Binary
// 2. Library/Package
// 3. Both

// 1. Each and every package must have a directory
// 2. Usually the name of the package and the name of the directory is same
// 	The name of the file not necessarly to the be the name of the directory or file name
