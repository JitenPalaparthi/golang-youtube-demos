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

func What() {
	println("Shapes is used to get area and perimeter for all geometric shapes")
}
