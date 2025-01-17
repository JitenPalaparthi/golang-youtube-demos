package main

import (
	"fmt"
	"reflect"
)

func main() {
	var arr1 [5]int // what is the type of this array?
	fmt.Println(arr1)

	arr1[0] = 10
	arr1[1] = 20
	arr1[2], arr1[3], arr1[4] = 30, 40, 50
	//fmt.Println(arr1)

	sum := 0 // creating a new variable
	for i := 0; i < len(arr1); i++ {
		sum += arr1[i]
		print(arr1[i], " ")
	}

	println("\nSum:", sum)

	sum = 0

	for _, v := range arr1 {
		sum += v
	}
	println("Sum:", sum)

	// creating an array using short hand declaration

	arr2 := [3]int{10, 20, 30}
	arr3 := [...]int{10, 20, 30, 40} // ... is evaluated as length at compile time

	fmt.Println("type of arr2", reflect.TypeOf(arr2))
	fmt.Printf("type of arr3: %T", arr3)

	s1 := SumOf(arr1)
	fmt.Println("Sumof arr1", s1)
	//SumOf(arr2)
	//SumOf(arr3)
	// arr1 and arr2 are two different types
}

// array has a fixed length. The length should be known or evaluated at compiletime

// zero value for array
// var arr1 [5]int --> zero value is [0 0 0 0 0]
// array index starts from 0
// the type of array includes its length [5]int --> type is [5]int
// cannot type cast arrays of different types

func SumOf(arr [5]int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}
