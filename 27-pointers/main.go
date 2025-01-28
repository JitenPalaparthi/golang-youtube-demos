package main

import (
	"errors"
	"fmt"
)

func main() {

	// const (
	// 	num int  = 100
	// 	ptr *int = &num
	// )

	// cannot create a constant pointer

	arr1 := [3]int{10, 11, 12}
	var ptr1 *[3]int // a pointer that points to the array

	ptr1 = &arr1 // contains the address of arr1

	(*ptr1)[0] = 100
	(*ptr1)[1] = 200
	ptr1[2] = 300 // without dereference operator..

	fmt.Println(*ptr1)
	slice1 := make([]int, 5, 10)
	for i := 0; i < len(slice1); i++ {
		slice1[i] = i + 1
	}
	AddElemsAndDouble(slice1, 6, 7, 8, 9, 10)
	fmt.Println("normal slice:", slice1)

	slice2 := make([]int, 5, 10)
	for i := 0; i < len(slice2); i++ {
		slice2[i] = i + 1
	}

	AddElemsAndDoubleP(&slice2, 6, 7, 8, 9, 10)
	fmt.Println("using pointer in the func:", slice2)

	// var num int = 100
	// incr(&num)

	var arr2 [3]*int  // an array has the length of 3 holding 3 addresses
	var arr3 *[3]*int // array that is holding the address of [3]* dint array

	a, b, c := 10, 20, 30
	arr2[0], arr2[1], arr2[2] = &a, &b, &c
	arr3 = &arr2

	fmt.Println(*arr3)

	fmt.Println("the address at the 0th element", *arr3[0])
	fmt.Println("the pointer arrays's 0th element", (*arr3)[0]) // no need to use to get the value of 0th element
	fmt.Println("the value arrays's 0th element", *(*arr3)[0])  // no need to use to get the value of 0th element

	//arr3[0] go can understand the above by calling it , just like a normal arr
	*arr3[0] = 9999
	fmt.Println("value of a:", a)

	// var slice3 []*int
	// var slice4 *[]int

}

func AddElemsAndDouble(slice []int, elems ...int) {
	slice = append(slice, elems...)
	for i, v := range slice {
		slice[i] = v * 2
	}
}

func AddElemsAndDoubleR(slice []int, elems ...int) []int {
	slice = append(slice, elems...)
	for i, v := range slice {
		slice[i] = v * 2
	}
	return slice
}

func AddElemsAndDoubleP(slice *[]int, elems ...int) error {
	if slice == nil {
		return errors.New("nil slice")
	}
	*slice = append(*slice, elems...)
	for i, v := range *slice {
		(*slice)[i] = v * 2
	}

	return nil
}

func incr(ptr *int) {
	*ptr++
}
