package main

import (
	"fmt"
	"unsafe"
)

func main() {

	slice1 := []int{100, 110, 120, 130, 140}
	// slice header
	// Ptr
	// Len
	// Cap

	ptr := uintptr(unsafe.Pointer(&slice1))

	header := (*[3]int)(unsafe.Pointer(ptr))
	fmt.Println("Ptr:", header[0])
	fmt.Println("Len:", header[1])
	fmt.Println("Cap:", header[2])

	header[1] = 100                                         // technically can change the len of slice but it is not a correct operation.Basically tampering the data
	fmt.Println("Len:", len(slice1), "\nCap:", cap(slice1)) // does not mean any sense when length is 100 but cap is still 5

	//

	var slice2 []int         // Is it nil?
	slice3 := make([]int, 0) // Is it nil?
	slice4 := []int{}        // Is it nil?
	slice5 := []int{}        // Is it nil?

	fmt.Println("slice2 internal details")

	header2 := (*[3]int)(unsafe.Pointer(uintptr(unsafe.Pointer(&slice2))))
	fmt.Println("slice2 Ptr:", header2[0]) // 0x0
	fmt.Println("slice2 Len:", header2[1]) // 0
	fmt.Println("slice2 Cap:", header2[2]) // 0
	println()
	fmt.Println("slice3 internal details")
	header3 := (*[3]int)(unsafe.Pointer(uintptr(unsafe.Pointer(&slice3))))
	fmt.Println("slice3 Ptr:", header3[0]) // some pointer
	fmt.Println("slice3 Len:", header3[1]) // 0
	fmt.Println("slice3 Cap:", header3[2]) // 0
	println()

	println()
	fmt.Println("slice4 internal details")
	header4 := (*[3]int)(unsafe.Pointer(uintptr(unsafe.Pointer(&slice4))))
	fmt.Println("slice4 Ptr:", header4[0]) // some pointer
	fmt.Println("slice4 Len:", header4[1]) // 0
	fmt.Println("slice4 Cap:", header4[2]) // 0
	println()

	println()
	fmt.Println("slice5 internal details")
	header5 := (*[3]int)(unsafe.Pointer(uintptr(unsafe.Pointer(&slice5))))
	fmt.Println("slice5 Ptr:", header5[0]) // some pointer
	fmt.Println("slice5 Len:", header5[1]) // 0
	fmt.Println("slice5 Cap:", header5[2]) // 0
	println()

	slice5 = append(slice5, 100)
	println()
	fmt.Println("slice5 internal details")
	header5 = (*[3]int)(unsafe.Pointer(uintptr(unsafe.Pointer(&slice5))))
	fmt.Println("slice5 Ptr:", header5[0]) // some pointer
	fmt.Println("slice5 Len:", header5[1]) // 0
	fmt.Println("slice5 Cap:", header5[2]) // 0
	println()

	var str1 string

	var str2 string

	var str3 string = "Hello World"

	// string header
	// pointer
	// length
	println("str1 header")
	strheader1 := (*[2]int)(unsafe.Pointer(uintptr(unsafe.Pointer(&str1))))
	fmt.Println("str1 ptr:", strheader1[0])
	fmt.Println("str1 len:", strheader1[1])
	println()

	println("str2 header")
	strheader2 := (*[2]int)(unsafe.Pointer(uintptr(unsafe.Pointer(&str2))))
	fmt.Println("str2 ptr:", strheader2[0])
	fmt.Println("str2 len:", strheader2[1])
	println()

	println("str3 header")
	strheader3 := (*[2]int)(unsafe.Pointer(uintptr(unsafe.Pointer(&str3))))
	fmt.Println("str3 ptr:", strheader3[0])
	fmt.Println("str3 len:", strheader3[1])
	println()

}
