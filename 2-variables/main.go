package main

import "fmt"

var Global int = 100 // package level variable , global variable

var (
	FirstName string = "John"
	LastName  string = "Doe"
)

const MAX int = 100
const (
	PI        float32 = 3.14
	SECOND            = 60
	MINITE            = 60 * SECOND // This is an constant expression
	HOUR              = 60 * MINITE
	DAY               = 24 * HOUR
	SOMETHING         = 100*200/MINITE + MAX%HOUR + 100
)

func main() {

	{ // block stats
		a, b := 10, 20 // assigning variables with short hand declaration
		println(a, b)  // printing a and b
	} // block ends , deallocate a and b from the stack

	// another example of runtime --> vtable in c++ itable (idt) in golang
	// runtime polymorphism --?
	var num1 int // on 64bit machine it is 8 bytes on 32bit machine it is 4 bytes

	const MIN int = 1
	var num2 uint8 = 123         // 1 byte
	var num3 uint32 = 23         // 4 bytes
	var flaot1 float32 = 12.34   // 4 bytes
	var float2 float64 = 123.123 // 8 bytes

	fmt.Println(num1)
	fmt.Println(num2)
	fmt.Println(num3)
	fmt.Println(flaot1)
	fmt.Println(float2)
	// declaring variables as a group
	var (
		num4 int
		num5 uint64
		ok1  bool
		str1 string
	)
	fmt.Println(num4, num5, ok1, str1)

	var num6 = 100 // declarting with type inference
	var float3 = 12.23
	var ok2 = true
	var str2 = "Hello UST Global"
	var age = 39 // age takes 8 bytes on 64bit machine and 4 bytes on 32bit machine
	fmt.Println(num6, float3, ok2, str2, age)

	// short hand declaration with type inference
	num7 := 12
	ok3 := true
	str3 := "Hello UST Global"
	fmt.Println(num7, ok3, str3)

	c1 := 12.2 + 2.0i // another way of declaring complex numbers
	var a, b float32 = 12.3, 1.3
	c2 := complex(a, b)      // very clear about the memory based on a and b
	c3 := complex(12.3, 1.3) // type inference works hence it takes as float64 and float64
	fmt.Println(c1, c2, c3)
}

// numbers
/*
 int,int8, int16, int32, int64
 uint, uint8, uint16, uint32, uint64
 float32, float64
 complex64, complex128
 uintptr
*/

// boolean --> bool
// strings --> string --> collection of bytes

// zero values -> default values -> for int 0 , for bool false, string "" for float 0.0
// type inference

// What is an expression?
// What is a statement?
// What is a constant expression?
// Where constants are stored --? in the symbol table
// Data segment and Code/Text segment
// What is runtime -->? runtime is a library which is written in go and it is used to run the go programs
// Runtime can be embedded in the executable or it can be linked dynamically
