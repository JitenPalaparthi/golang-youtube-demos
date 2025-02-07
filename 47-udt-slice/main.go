package main

import "fmt"

func main() {

	slice1 := make([]int, 5)
	slice1[0], slice1[1], slice1[2], slice1[3], slice1[4] = 10, 11, 12, 13, 14
	sum1 := IntSlice(slice1).GetSum()
	fmt.Println(sum1)

	slice2 := make(IntSlice, 5)
	slice2[0], slice2[1], slice2[2], slice2[3], slice2[4] = 10, 11, 12, 13, 14

	sum2 := slice2.GetSum()
	fmt.Println(sum2)

	// slice3 := make([]int64, 5)
	// slice3[0], slice3[1], slice3[2], slice3[3], slice3[4] = 10, 11, 12, 13, 14
	// IntSlice(slice3).GetSum() // This does not work

}

type IntSlice []int

func (is IntSlice) GetSum() int {
	sum := 0
	for _, v := range is {
		sum += v
	}
	return sum
}
