package main

import "fmt"

func main() {

	// var any1 any = 100
	// dataptr is pointer to 100
	// typeptr is pointer to a type called int

	r1 := new(Rect)
	r1.L, r1.B = 12.3, 12.23

	var s1 Square = 12.34

	var iarea1 IArea = r1 // This is qualified

	fmt.Println("Area of r1:", iarea1.Area())
	iarea1.Display()

	iarea1 = s1
	iarea1.Display()
	fmt.Println("Area of s1:", iarea1.Area())

	var ishape1 IShape
	ishape1 = r1
	fmt.Println("Area of r1:", ishape1.Area())

}

type IArea interface {
	Area() float32
	Display()
}
type IShape interface {
	Area() float32
}

type Rect struct {
	L, B float32
}

func (r *Rect) Area() float32 {
	return r.L * r.B
}

func (r *Rect) Display() {
	fmt.Println("Rect")
}

type Square float32

func (s Square) Area() float32 {
	return float32(s * s)
}

func (r Square) Display() {
	fmt.Println("Square")
}

// IArea is not different than any / interface{}
// any header --> DataPtr , TypePtr
// Interface Definition Table , iTable (VTable in c++, Rust)
