package main

import "fmt"

func main() {

	r1 := NewRect(12.3, 34.3)
	a1 := r1.Area()
	p1 := r1.Perimeter()

	fmt.Println("Area of rect r1:", a1)
	fmt.Println("Perimter of rect r1:", p1)

	s1 := NewSquare(12.3)
	a2 := s1.Area()
	p2 := s1.Perimeter()
	fmt.Println("Area of square s1:", a2)
	fmt.Println("Perimter of square r1:", p2)

	var ishape IShape
	ishape = r1

	fmt.Println("Area of", ishape.Display(), ":", ishape.Area())
	fmt.Println("Perimeter of", ishape.Display(), ":", ishape.Perimeter())
	println()

	ishape = s1
	fmt.Println("Area of", ishape.Display(), ":", ishape.Area())
	fmt.Println("Perimeter of", ishape.Display(), ":", ishape.Perimeter())
	println()

}

func Shape(ishape IShape) {
	fmt.Println("Area of", ishape.Display(), ":", ishape.Area())
	fmt.Println("Perimeter of", ishape.Display(), ":", ishape.Perimeter())
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

// Static Dispatch vs Dynamic Dispatch
