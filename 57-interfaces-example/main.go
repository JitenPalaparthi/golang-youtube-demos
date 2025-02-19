package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var ishape IShape
	if ishape == nil {
		println("Yes it is nil")
	} // Both ptrs are nil

	// a := ishape.Area()
	// fmt.Println(a)

	var c Circle = 123.23

	ishape = c // ?

	a1 := c.Area()
	fmt.Println("Area a1", a1)

	a := ishape.Area()
	fmt.Println(a)

	var s Square = 123.123
	// ishape = Square(123.123)
	ishape = s
	// DataPtr
	// TypePtr is Square

	a = ishape.Area()
	fmt.Println(a)

	ptr := (*[2]unsafe.Pointer)(unsafe.Pointer(uintptr(unsafe.Pointer(&ishape))))
	fmt.Println((ptr)[0])
	fmt.Println((ptr)[1])

	fmt.Println((*float32)(ptr[1]))
	//fmt.Printf("0X%x\n", &s)

	ishape = &Rect{L: 12.23, B: 343.34}

}

type IShape interface {
	Area() float32
}

const PI = 3.14

type Circle float32
type Square float32

type Rect struct {
	L, B float32
}

func (r *Rect) Area() float32 {
	return r.L * r.B
}

func (c Circle) Area() float32 {
	return float32(PI * c * c)
}

func (s Square) Area() float32 {
	return float32(s * s)
}

// DataPtr --> What data that it has pointed to
// TypePtr --> What concrete type it has pointed to
