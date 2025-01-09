package main

import "fmt"

func main() {

	var num1 uint8 = 123
	//var num2 int = num1 // This is not allowed in go even num1 is 1 byte and jnum2 is 8 bytes on 64bit
	var num2 = int(num1)
	fmt.Println(num2)

	// What all types can be cased ?
	// all number types can be casted to each other

	// var ok1 bool = true

	// var num3 uint8 = uint8(ok1) // a bool cannot be type casted to number
	float1 := 123.34 // float64

	var num3 int = int(float1)
	println(num3)
}

// Go there is not implicit typecasting

// typecasting is a way to convert a variable from one type to another type
