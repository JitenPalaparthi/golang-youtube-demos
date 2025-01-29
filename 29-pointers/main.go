package main

import (
	"fmt"
	"unsafe"
)

func main() {

	arr := [5]int{10, 11, 12, 13, 14}

	// for _, v := range arr {
	// 	println(v)
	// }

	uptr := uintptr(unsafe.Pointer(&arr[0])) // 1 and 4
	//fmt.Println(uptr)
	//fmt.Printf("0x%x\n", uptr)
	uptr += 8 //unsafe.Sizeof(arr[0])

	v := (*int)(unsafe.Pointer(uptr)) // 3 2
	fmt.Println(*v)
	*v = 1111

	fmt.Println(arr)

	fmt.Println("perform iteration using unsafe.Pointer")
	// iterate thru the array in the same way

	ptr := uintptr(unsafe.Pointer(&arr[0]))
	for i := 0; i < len(arr); i++ {
		v := (*int)(unsafe.Pointer(ptr)) // 3 2
		fmt.Println(*v)
		ptr += 8
	}

}

// unsafe.Pointer()
// 1. A pointer value of any type can be converted to a Pointer.
// 2. A Pointer can be converted to a pointer value of any type.
// 3. A uintptr can be converted to a Pointer.
// 4. A Pointer can be converted to a uintptr.
