package main

import (
	"fmt"

	. "github.com/JitenPalaparthi/shapes1/shape"
	rectangle "github.com/JitenPalaparthi/shapes1/shape/rect"
	"github.com/JitenPalaparthi/shapes1/shape/square"
	"github.com/google/uuid"
)

func main() {
	r1 := rectangle.NewRect(12.4, 34.3)
	s1 := square.Square(12.2)
	//var ishape shape.IShape
	Shape(r1)
	Shape(s1)
	What()
	uuid := uuid.New()
	fmt.Println(uuid.String())
}
