package rectangle

import "fmt"

// var (
// private string
// Public  string
// )
//
// type t struct { // restricted
// Name string // unrestricted
// }

func (r *Rect) display() {
	// cannot call display in another package due to it is unexported
	fmt.Println("Rectangle")
	fmt.Println("L:", r.L)
	fmt.Println("B:", r.B)
}

type Rect struct {
	L, B float32 // unrestricted fileds
	a, p float32 // restricted fields
}

func NewRect(l, b float32) *Rect {
	return &Rect{L: l, B: b}
}

func (r *Rect) Display() {
	r.display() // this can be called locally inside a pacakge
}

func (s *Rect) What() string {
	return "Rectangle"
}

// Restructed or unrestricted
// Exported or unexported
// Any thing(a package level variable, or a type or a function) that stats with
// Uppercase is exportable, if it starts with lowerCase it is unexportable
// UpperCase is similar to public
// LowerCase is similar to private
