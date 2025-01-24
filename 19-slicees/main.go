package main

import "fmt"

func main() {

	slice1 := []int{10, 20, 30, 40, 50}
	// Ptr, Len:5 , Cap:5
	fmt.Println("slice1:", slice1)
	fmt.Printf("Address of the slice:%p\n", &slice1)
	fmt.Println("Len:", len(slice1), "\nCap:", cap(slice1))
	fmt.Printf("Ptr of the slice:%p\n", &slice1[0])

	println()
	slice1 = append(slice1, 9999)
	fmt.Println("slice1:", slice1)
	fmt.Printf("Address of the slice:%p\n", &slice1)
	fmt.Println("Len:", len(slice1), "\nCap:", cap(slice1))
	fmt.Printf("Ptr of the slice:%p\n", &slice1[0])

	println()
	slice1 = append(slice1, 8888)
	fmt.Println("slice1:", slice1)
	fmt.Printf("Address of the slice:%p\n", &slice1)
	fmt.Println("Len:", len(slice1), "\nCap:", cap(slice1))
	fmt.Printf("Ptr of the slice:%p\n", &slice1[0])

	println()
	slice1 = append(slice1, 7777, 6666, 5555, 4444)
	fmt.Println("slice1:", slice1)
	fmt.Printf("Address of the slice:%p\n", &slice1)
	fmt.Println("Len:", len(slice1), "\nCap:", cap(slice1))
	fmt.Printf("Ptr of the slice:%p\n", &slice1[0])

}
