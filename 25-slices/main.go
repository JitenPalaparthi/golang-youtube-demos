package main

import (
	"fmt"
)

func main() {

	var slice1 []int
	slice2 := make([]int, 0)
	slice3 := []int{} // slice declartion with empty literal

	if slice1 == nil {
		println("slice1 is nil")
	}

	if slice2 == nil {
		println("slcie2 is nil")
	}

	if slice3 == nil {
		println("slcie3 is nil")
	}

	slice4 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// to delete an element from the slice

	slice5 := append(slice4[:3], slice4[4:]...)
	//slice9 := append(slice4[:3], 5, 6, 7, 8, 9, 10)
	fmt.Println(slice5)
	// deep copy
	slice6 := make([]int, 10)
	copy(slice6, slice5)

	fmt.Println("slcie6", slice6)

	// copy built in function does the below
	slice7 := make([]int, 10)
	for i := 0; i < min(len(slice7), len(slice5)); i++ {
		slice7[i] = slice5[i]
	}

	fmt.Println("slcie7", slice7)

	slice8 := make([]int, 4)
	copy(slice8, slice5) // only 4 elements gets copied
	fmt.Println("slcie8", slice8)

	clear(slice8) // clears all the elemens to zero values
	fmt.Println("slcie8", slice8)

	// slice9 := make([]int, 0)
	// fmt.Printf("slice9: %+v, Is nil: %t\n", slice9, slice9 == nil)

}

// create a function to delete elements from the slice

func DeleteElem(slice []int) error {
	if slice == nil {
		//return errors.New("nil slice")
		return fmt.Errorf("nil slice")
	}
	return nil
}

func Calc(a, b int) (int, int) {
	return a + b, a - b
}
