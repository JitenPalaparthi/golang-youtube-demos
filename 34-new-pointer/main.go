package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var num int = 12321
	var ptr1 *int = &num
	fmt.Println(*ptr1)

	var ptr2 *int = new(int)
	// new allocates 8 bytes (on 64 bit machines) and assign zero value to that address
	fmt.Println(*ptr2)
	*ptr2 = 1000
	fmt.Println(*ptr2)
	var ptr3 *bool = new(bool)
	fmt.Println(*ptr3)

	var ptr4 *string = new(string)
	fmt.Println(ptr4)

	var arrPtr *[5]int = new([5]int)
	arrPtr[0] = 100
	arrPtr[1] = 110

	for _, v := range *arrPtr {
		println(v)
	}

	var slicePtr *[]int = new([]int)
	if *slicePtr == nil {
		fmt.Println("nil slice pointer")
	}
	fmt.Println("Len:", len(*slicePtr), "Cap:", cap(*slicePtr))

	//ptr7 := (*[3]int)(unsafe.Pointer(uintptr(unsafe.Pointer(slicePtr))))
	ptr7 := (*[3]int)(unsafe.Pointer(slicePtr))
	//ptr8 := (*[3]int)(unsafe.Pointer(slicePtr))
	fmt.Println("Address:", ptr7[0], "Len:", ptr7[1], "Cap:", ptr7[2])
	*slicePtr = append(*slicePtr, 100)
	fmt.Println("Address:", ptr7[0], "Len:", ptr7[1], "Cap:", ptr7[2])

	//m1 := map[string]any{"hi": "Hello"}
	m2 := new(map[string]any)

	c := *(*[5]int)(unsafe.Pointer(&m2))
	fmt.Println(c)

}
