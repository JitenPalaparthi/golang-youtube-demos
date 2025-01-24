package main

import (
	"fmt"
	"math"
)

func main() {

	slice1 := []int{10, 20, 30, 40, 50}
	// Ptr, Len:5 , Cap:5
	slice1 = append(slice1, 60)
	fmt.Println("slice1:", slice1)
	fmt.Printf("Address of the slice1:%p\n", &slice1)
	fmt.Println("Len:", len(slice1), "\nCap:", cap(slice1))
	fmt.Printf("Ptr of the slice1:%p\n", &slice1[0])
	sqSlice(slice1)
	println()
	fmt.Println("slice1:", slice1)
	fmt.Printf("Address of the slice1:%p\n", &slice1)
	fmt.Println("Len:", len(slice1), "\nCap:", cap(slice1))
	fmt.Printf("Ptr of the slice1:%p\n", &slice1[0])

	sqRootSlice(slice1)
	println()
	fmt.Println("slice1:", slice1)

	slice1 = addSqSliceR(slice1, 70)
	//addSqSlice(slice1, 70)
	println()

	fmt.Println("slice1:", slice1)
	fmt.Printf("Address of the slice1:%p\n", &slice1)
	fmt.Println("Len:", len(slice1), "\nCap:", cap(slice1))
	fmt.Printf("Ptr of the slice1:%p\n", &slice1[0])

}

func sqSlice(slice []int) { // pass or call by value
	for i, v := range slice {
		slice[i] = v * v
	}
}

func sqRootSlice(slice []int) { // pass or call by value
	for i, v := range slice {
		slice[i] = int(math.Sqrt(float64(v)))
	}
}

func addSqSliceR(slice []int, elem int) []int { // pass or call by value
	slice = append(slice, elem) // adding a new element to the slice and then sq the elements of the slice
	for i, v := range slice {
		slice[i] = v * v
	}
	return slice
}

func addSqSlice(slice []int, elem int) { // pass or call by value
	slice = append(slice, elem) // adding a new element to the slice and then sq the elements of the slice
	for i, v := range slice {
		slice[i] = v * v
	}
}
