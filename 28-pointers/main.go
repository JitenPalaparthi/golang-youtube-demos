package main

import (
	"fmt"
)

func main() {

	arr1 := [5]int{10, 11, 12, 13, 14}
	// cgo
	fmt.Println(&arr1[0])

	//var intPtr uintptr

	var num = 100
	var ptr = &num
	//ptr += 8 // This is called pointer arithmetic, this is not allowed in go

	fmt.Printf("address ptr has: %p\naddress ptr has: %d\naddress ptr has: %b\n", ptr, ptr, ptr)

	// unsafe.Pointer()
	// 1. A pointer value of any type can be converted to a Pointer.
	// 2. A Pointer can be converted to a pointer value of any type.
	// 3. A uintptr can be converted to a Pointer.
	// 4. A Pointer can be converted to a uintptr.

}
