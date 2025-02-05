package main

import "fmt"

func main() {

	r1 := Rect{L: 12.4, B: 13.64}
	a1 := Area(&r1)
	p1 := Perimeter(&r1)

	fmt.Printf("Area a1 of r1:%.2f\n", a1)
	fmt.Printf("Perimeter p1 of r1:%.2f\n", p1)

	fmt.Println("Printing direct values of Area and Perimeter of r1:", r1.A, r1.P)

	r2 := NewRect(12.4, 13.64)

	a2, p2 := Area(r2), Perimeter(r2)
	fmt.Println("Area a2 of r2:", a2, "\nPerimeter p2 of r2:", p2)
	fmt.Println("Printing direct values of Area and Perimeter of r2:", r2.A, r2.P)

	r3 := NewRect(12.4, 13.64)
	a3, p3 := r3.Area(), r3.Perimeter()

	fmt.Println("Area a3 of r3:", a3, "\nPerimeter p3 of r3:", p3)
	fmt.Println("Printing direct values of Area and Perimeter of r3:", r3.A, r3.P)

	r4 := new(Rect)
	r4.L = 12.4
	r4.B = 13.64

	a4, p4 := r4.Area(), r4.Perimeter()

	fmt.Println("Area a4 of r4:", a4, "\nPerimeter p4 of r4:", p4)
	fmt.Println("Printing direct values of Area and Perimeter of r4:", r4.A, r4.P)

}

type Rect struct {
	L, B float32
	A, P float32
}

func NewRect(l, b float32) *Rect {
	return &Rect{L: l, B: b}
}

// Area function of a Rect
func Area(r *Rect) float32 {
	// (*r).A = r.L * r.B
	r.A = r.L * r.B
	return r.A
}

// Perimeter function of a Rect
func Perimeter(r *Rect) float32 {
	r.P = 2 * (r.L + r.B)
	return r.P
}

func (r *Rect) Area() float32 {
	// (*r).A = r.L * r.B
	r.A = r.L * r.B
	return r.A
}

func (r *Rect) Perimeter() float32 {
	r.P = 2 * (r.L + r.B)
	return r.P
}
