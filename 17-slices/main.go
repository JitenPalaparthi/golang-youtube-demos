package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var slice1 []int // slice declaration but not instantiation
	if slice1 == nil {
		println("slice1 is a nil slice")
	}
	fmt.Println("Size of the slice", unsafe.Sizeof(slice1))
	//fmt.Println(slice1[0]) // indexout of bound error
	slice1 = make([]int, 5) // allocate memory to the slice with length 5.

	slice2 := make([]int, 5, 10) // instantiate with 5 length and 10 is the cap

	slice3 := []int{10, 12, 13, 14, 15}
	arr1 := [5]int{10, 12, 13, 14, 15}
	arr2 := [...]int{10, 12, 13, 14, 15}

	fmt.Printf("Address of the slice1 header:%p Len:%d Cap:%d, Ptr:%p,Actual data:%v", &slice1, len(slice1), cap(slice1), &slice1[0], slice1)
	/*
		slice1
			Ptr: 0x123
			Len: 5
			Cap: 5
	*/
	fmt.Println(slice2) //
	/*
		slice2
			Ptr: 0xff11
			Len: 5
			Cap: 10
	*/
	fmt.Println(slice3)
	/*
		slice3
			Ptr: 0x1122
			Len: 5
			Cap: 5
	*/
}

/* sliceheader
Ptr  -> The pointer to the actual data, 8 bytes
Len  -> Len of the slice , 8 bytes
Cap  -> Cap of the slice , 8 bytes
*/
// 24 bytes
// make is a built in function
// make is used to instantiate the slice, map and also channel
// zero values to the slice, based on the length
