package main

import "fmt"

func main() {
	slice1 := []int{10, 20, 30, 40, 50}
	// Ptr, Len:5 , Cap:5
	fmt.Println("slice1:", slice1)
	fmt.Printf("Address of the slice1:%p\n", &slice1)
	fmt.Println("Len:", len(slice1), "\nCap:", cap(slice1))
	fmt.Printf("Ptr of the slice1:%p\n", &slice1[0])

	slice2 := slice1
	// slice1 is a structure, which is header
	// slice2 also has a structure the actual data is the header
	println()
	fmt.Println("slice2:", slice2)
	fmt.Printf("Address of the slice2:%p\n", &slice2)
	fmt.Println("Len:", len(slice2), "\nCap:", cap(slice2))
	fmt.Printf("Ptr of the slice2:%p\n", &slice2[0])

	slice2[1] = 9999
	fmt.Println("slice1", slice1)
	fmt.Println("slice2", slice2)

	slice2 = append(slice2, 8888)
	slice2[0] = 1111

	//slice1 = slice2 // headers are copied
	// fmt.Println("slice1", slice1)
	// fmt.Println("slice2", slice2)
	//slice2[0] = 1111
	println()
	fmt.Println("slice1:", slice1)
	fmt.Printf("Address of the slice1:%p\n", &slice1)
	fmt.Println("Len:", len(slice1), "\nCap:", cap(slice1))
	fmt.Printf("Ptr of the slice1:%p\n", &slice1[0])

	println()
	fmt.Println("slice2:", slice2)
	fmt.Printf("Address of the slice2:%p\n", &slice2)
	fmt.Println("Len:", len(slice2), "\nCap:", cap(slice2))
	fmt.Printf("Ptr of the slice2:%p\n", &slice2[0])

}

// slice header
// Ptr
// Len
// Cap
