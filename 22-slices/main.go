package main

import "fmt"

func main() {

	arr1 := [5]int{10, 11, 12, 13, 14}
	sum1 := SumOfArr(arr1)
	fmt.Println("Sum of elements of arr1:", sum1)

	arr2 := [4]int{10, 11, 12, 13}

	//sum2 := SumOfArr(([4]int)(arr2))
	sum2 := SumOf(arr2[:])
	sum3 := SumOf(arr1[:])
	fmt.Println("Sum of elements of arr2:", sum2)
	fmt.Println("Sum of elements of arr1:", sum3)

	slice1 := arr2[:]
	// array does not have the header but when use arr[:] it becomes slice
	// and it has the header
	// slice1
	// Ptr , Len, Cap
	// arr2 : &ptr[0] and len
	sum4 := SumOf(slice1)
	fmt.Println("Sum of elements of slice1:", sum4)
	slice1[0] = 9999
	fmt.Println(arr2)
	fmt.Println("slice1 cap", cap(slice1))
	slice1 = append(slice1, 100) // since the cap is 4 the len has to be 5 , it reallocates the slice
	// and slice and arr are divergent
	slice1[1] = 8888
	fmt.Println(arr2)
	fmt.Println("slice1 cap", cap(slice1))

}

func SumOf(slice []int) int {
	sum := 0
	for _, v := range slice {
		sum += v
	}
	return sum
}

func SumOfArr(arr [5]int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}
