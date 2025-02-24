package shape

import "fmt"

type IShape interface {
	Area() float32
	Perimeter() float32
	What() string
}

func Shape(ishape IShape) {
	fmt.Println("Area of ", ishape.What(), ":", ishape.Area())
	fmt.Println("Perimeter of ", ishape.What(), ":", ishape.Perimeter())
}

func Display() {
	println("shape package contains all shapes")
}
