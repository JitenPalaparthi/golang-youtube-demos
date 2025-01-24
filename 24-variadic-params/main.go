package main

import "fmt"

func main() {
	fmt.Println("Hello", "World", true, 100, (100 / 2), true && false, 1231.123)
	slice1 := []int{} //
	//slice2 := make([]int, 0)
	slice1 = append(slice1, 10, 11, 12, 123, 1, 34, 35, 5, 450, 44, 76, 9, 8, 34, 57)
	fmt.Println(slice1)

	fmt.Println(SumOf())
	fmt.Println(SumOf(10, 12))
	fmt.Println(SumOf(10, 12, 123, 213))
	fmt.Println(SumOf(10, 11, 12, 123, 1, 34, 35, 5, 450, 44, 76, 9, 8, 34, 57))
	s1 := SumOf(slice1...) // slice can be
	arr1 := [14]int{10, 11, 12, 123, 1, 34, 35, 5, 450, 44, 76, 9, 8, 34}
	s2 := SumOf(arr1[:]...) // array  can be after converting it to slice
	fmt.Println(s1, s2)
}

// variadic parameter ...int
// variadic paramter must be the last parameter in a func
// variadic paramter cannot be used any where else other than function or method input parameters
// eclipse symbol is used to denote variadic parameters
// to iterate , use range loop
func SumOf(nums ...int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum
}

// the variadic param must be the last param..
// func MultiplySumOf(nums ...int, m int) int {
func MultiplySumOf(m int, nums ...int) int {
	sum := 0
	for _, v := range nums {
		sum += v * m
	}
	return sum
}
