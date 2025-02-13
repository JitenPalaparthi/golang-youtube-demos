package main

import "fmt"

func main() {

	//var num int = 100
	var ishape IShape = NewRect(12.4, 34.34)
	ishape = NewSquare(100)

	if ishape == nil {
		fmt.Println("yes nil")
	}

	shapes := make([]IShape, 0)
	// shapes := make([]any, 0)

	shapes = append(shapes, &Rect{L: 10.5, B: 12.5}, NewRect(10.1, 12.4), NewRect(12.5, 13.5), NewSquare(100.23), NewRect(45.6, 76.5), NewSquare(10.5), NewCuboid(12.4, 5.6, 7.8))

	for _, v := range shapes {
		Shape(v)
		//Shape(v.(IShape))

		//c, ok := (v.(IShape)).(*Cuboid)

		c, ok := v.(*Cuboid)
		if ok {
			fmt.Println("Calling a special method on a concrete type Cuboid")
			c.What()
		}

	}

}

func Shape(ishape IShape) {
	fmt.Println("Area of", ishape.Display(), ":", ishape.Area())
	fmt.Println("Perimter of", ishape.Display(), ":", ishape.Perimeter())
	println()
}

// Dependency Injection contains
// A dependent object and an independent Object

type IShape interface {
	Area() float32
	Perimeter() float32
	IDisplay
}

type IDisplay interface {
	Display() string
}

type Rect struct {
	L, B float32
}

func NewRect(l, b float32) *Rect {
	return &Rect{L: l, B: b}
}

func (r *Rect) Area() float32 {
	return r.L * r.B
}

func (r *Rect) Perimeter() float32 {
	return 2 * (r.L + r.B)
}

func (r *Rect) Display() string {
	return "Rect"
}

type Square float32

func NewSquare(s float32) Square {
	return Square(s)
}

func (s Square) Area() float32 {
	return float32(s * s)
}

func (s Square) Perimeter() float32 {
	return float32(4 * s)
}

func (r Square) Display() string {
	return "Square"
}

type Cuboid struct {
	L, B, H float32
}

func NewCuboid(l, b, h float32) *Cuboid {
	return &Cuboid{L: l, B: b, H: h}
}

func (c *Cuboid) Area() float32 {
	return c.L * c.B * c.H
}
func (c *Cuboid) Perimeter() float32 {
	return 2 * (c.L + c.B + c.H)
}

func (c *Cuboid) Display() string {
	return "Cuboid"
}

func (c *Cuboid) What() {
	fmt.Println("This is Cuboid Shape")
}
