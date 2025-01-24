package main

import (
	"fmt"
	v2 "math/rand/v2"
)

func main() {

	str1 := "Hello World"

	fmt.Printf("Address of str1:%p Ptr of the str1:%p\n", &str1, &([]byte(str1))[0])

	str1 = "Hello Universe" // This is called as mutability
	fmt.Printf("Address of str1:%p Ptr of the str1:%p\n", &str1, &([]byte(str1))[0])

	str2 := str1 // every time a new string is reassigned from an existing string, it allocates memory

	fmt.Printf("Address of str2:%p Ptr of the str2:%p\n", &str2, &([]byte(str2))[0])

	// str Header
	// Ptr --> Points to the actual data
	// Len --> Len of actual data

	slice1 := make([]int, 10, 20)

	//var slice2 []int // nil slice, this slice has been decleared but not instantiated
	PrintSliceHeader(slice1, "slice1")

	// println("slice1")
	// fmt.Printf("Address of Slice Header:%p\nLen:%d\nCap:%d\n", &slice1, len(slice1), cap(slice1))
	// if slice1 != nil {
	// 	fmt.Printf("Ptr:%p\n", &slice1[0])
	// 	fmt.Println(slice1)
	// }
	// println("-----------------------------")

	// for i := range slice1 {
	// 	slice1[i] = v2.IntN((len(slice1) + 1) * 10)
	// }
	FillSliceWithRand(slice1)
	PrintSliceHeader(slice1, "slice1")

	slice1 = append(slice1, 9999) // to append an existing slice
	PrintSliceHeader(slice1, "slice1")

	// println("slice1")
	// fmt.Printf("Address of Slice Header:%p\nLen:%d\nCap:%d\n", &slice1, len(slice1), cap(slice1))
	// if slice1 != nil {
	// 	fmt.Printf("Ptr:%p\n", &slice1[0])
	// 	fmt.Println(slice1)
	// }
	// println("-----------------------------")

	//PrintSliceHeader(slice2, "slice2")

	// println("slice2")
	// fmt.Printf("Address of Slice Header:%p\nLen:%d\nCap:%d\n", &slice2, len(slice2), cap(slice2))
	// if slice2 != nil {
	// 	fmt.Printf("Ptr:%p\n", &slice2[0])
	// 	fmt.Println(slice2)
	// }
	// println("-----------------------------")

}

// append is a built in function, to add new elements to the slice
//

func PrintSliceHeader(slice []int, sliceName string) {
	println(sliceName)
	fmt.Printf("Address of Slice Header:%p\nLen:%d\nCap:%d\n", &slice, len(slice), cap(slice))
	if slice != nil {
		fmt.Printf("Ptr:%p\n", &slice[0])
		fmt.Println(slice)
	}
	println("-----------------------------")
}

func FillSliceWithRand(slice []int) []int {
	for i := range slice {
		slice[i] = v2.IntN((len(slice) + 1) * 10)
	}
	return slice
}
