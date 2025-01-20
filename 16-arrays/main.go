package main

import (
	"fmt"
	"math/rand"
)

func main() {

	var arr1 [15]int

	for i := 0; i < len(arr1); i++ {
		arr1[i] = rand.Intn(100)
	}

	fmt.Println(arr1)

	var arr2d [2][2]int

	for i1, a1 := range arr2d {
		for i2 := range a1 {
			arr2d[i1][i2] = rand.Intn(100)
		}
	}

	//fmt.Println(arr2d)

	for _, a1 := range arr2d {
		for _, v := range a1 {
			print(v, " ")
		}
		println()
	}

	arr2d1 := [2][2]int{{1, 2}, {3, 4}}
	fmt.Println(arr2d1)

	arr3d := [2][2][2]int{{{1, 2}, {3, 4}}, {{5, 6}, {7, 8}}}
	println("print 3darr")

	for _, a1 := range arr3d {
		for _, a2 := range a1 {
			for _, v := range a2 {
				print(v, " ")
			}
		}
	}
	println()

	// str1 := "Hello World"
	// var any1 any = 123.23
}

/*
String Header
Ptr -> 8 bytes
Len -> 8 bytes
*/

/*
any header
DataPtr --> Data 8 bytes
TypePtr --> Type information 8 bytes
*/
