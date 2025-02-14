package main

import "fmt"

func main() {

	r1 := NewRect(10.2, 12.4)

	r1.Display()
	r1.SetBorder(10)
	r1.SetBGColour("Red")
	r1.Display()
	r2 := r1.SetBorderColour("Green").(*Rect)

	r2.Display()

	r3 := NewRect(12.34, 14.56).SetBGColour("Orange").SetBorder(2).SetBorderColour("Blue").(*Rect)

	r3.Display()
}

type Rect struct {
	L, B         float32
	Border       float32
	BorderColour string
	BGColour     string
}

func NewRect(l, b float32) *Rect {
	return &Rect{L: l, B: b}
}

func (r *Rect) SetBGColour(c string) IShape {
	r.BGColour = c
	return r
}

func (r *Rect) SetBorder(b float32) IShape {
	r.Border = b
	return r
}

func (r *Rect) SetBorderColour(c string) IShape {
	r.BorderColour = c
	return r
}

func (r *Rect) Display() {
	fmt.Println("Length:", r.L, "Brewdth:", r.B, "Border:", r.Border, "Border Colour:", r.BorderColour, "BackGround Colour:", r.BGColour)
}

type IShape interface {
	SetBorder(b float32) IShape
	SetBorderColour(colour string) IShape
	SetBGColour(colour string) IShape
}
