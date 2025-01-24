package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// fmt.Printf("Address of Slice Header:%p\nLen:%d\nCap:%d\n", &slice1, len(slice1), cap(slice1))
	// println()
	PrintSliceHeader(slice1, "slice1")

	slice2 := slice1[:] //slice1
	PrintSliceHeader(slice2, "slice2")

	slice3 := slice1[:5] // but not 5 0-4
	PrintSliceHeader(slice3, "slice3")

	slice4 := slice1[2:8] // 2nd index
	PrintSliceHeader(slice4, "slice4")
	slice5 := slice1[5:]
	PrintSliceHeader(slice5, "slice5")
}

// starting point is the index , and ending point is not ==
// 2:8 -> 2 is the index and upto 8 but not 8th index
func PrintSliceHeader(slice []int, sliceName string) {
	println(sliceName)
	fmt.Printf("Address of Slice Header:%p\nLen:%d\nCap:%d\n", &slice, len(slice), cap(slice))
	if slice != nil {
		fmt.Printf("Ptr:%p\n", &slice[0])
		fmt.Println(slice)
	}
	println("-----------------------------")
}
