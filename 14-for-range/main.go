package main

import "fmt"

func main() {

	str1 := "Hello World"
	str2 := "你好, UST Global"
	fmt.Println(str1, "\n", str2)
	// what is a string?

	for i := 0; i < len(str1); i++ {
		print(string(str1[i]), " ")
	}
	println()
	for i := 0; i < len(str2); i++ {
		print(string(str2[i]), " ")
	}
	println()

	// there is a loop to work with collections
	// strings, arrays, slices, maps , channels , range of values
	// the return values of range loop differs based on what type it is being used

	for i, v := range str2 {
		println(i, "-->", string(v))
	}

	println()
	for _, v := range str1 {
		println(string(v))
	}

	// println()
	// for i := range str1 {
	// 	println(i)
	// }

	//i := 1
	//for i := 0; i < 10; i++
	for i := range 10 { // new one from version 1.22 onwards
		print(i, " ")
	}

	a1, p1 := Rect(10.23, 9.45)
	fmt.Println("\nArea a1:", a1, "perimeter:", p1)
	a2, _ := Rect(12.23, 13.45) // _ blank identifier
	fmt.Println("Area a2:", a2)

	_, p2 := Rect(10.23, 9.45)
	fmt.Println("perimeter:", p2)
}

func Rect(length, breadth float32) (area float32, perimeter float32) {
	area, perimeter = length*breadth, 2*(length+breadth)
	return area, perimeter
}
