package main

import (
	"fmt"
	"reflect"
)

func main() {

	//var any1 interface{} // What is zero value of any?
	var any1 any // any is statically types variable
	if any1 == nil {
		println("yes it is nil")
	}

	fmt.Println("Value of any:", any1, "Type of any:", reflect.TypeOf(any1))

	any1 = 100
	fmt.Println("Value of any:", any1, "Type of any:", reflect.TypeOf(any1))

	any1 = "Hello World"
	fmt.Println("Value of any:", any1, "Type of any:", reflect.TypeOf(any1))

	any1 = true
	fmt.Println("Value of any:", any1, "Type of any:", reflect.TypeOf(any1))

	any1 = 1231.43
	fmt.Println("Value of any:", any1, "Type of any:", reflect.TypeOf(any1))

	any1 = 'A' // int32 or rune
	fmt.Println("Value of any:", any1, "Type of any:", reflect.TypeOf(any1))

	any1 = 1232.123 // it is type float

	//var float1 float64 = float64(float64) // this does not work, cannot do the type casting
	var float1 float64 = any1.(float64) // type assertion
	fmt.Println(float1)

	// var num2 int = any1.(int) // this does not work becase ,the type in any is float64 not int, so you cannot assert to int
	// fmt.Println(num2)

	num2, ok2 := any1.(int)
	if ok2 {
		println(num2)
	} else {
		println("unable to assert to int")
	}

}

/*
any is an empty interface
any type or any kind of data satisfies the empty interface
zero value of any is nil
can check whether any variable is nil or not]
type casting cannot be done on any, there is a concept called type assertion
any.(t)
any header
---------
data ptr -> pointer to the actual data
type ptr -> pointer to the type, all types are loaded and that type pointer is given here
*/
