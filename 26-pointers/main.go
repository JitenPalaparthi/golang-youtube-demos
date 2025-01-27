package main

import "fmt"

func main() {

	var num1 int = 123
	var ptr1 *int = &num1

	var ptr2 *int
	fmt.Println(ptr2) // what is the zero value of a pointer, nil

	fmt.Println("num1:", *ptr1)

	*ptr1 = 9999 // dereference
	fmt.Println("num1:", num1)
	Incr(num1)
	fmt.Println("num1:", num1)
	IncrP(&num1)
	fmt.Println("num1:", num1)
	IncrP(ptr1)
	fmt.Println("num1:", num1)
	//num1 = 8888
	//*(&num1) = 7777
	//fmt.Println("num1:", num1)

	var ptr3 **int
	fmt.Println(ptr3)
	ptr3 = &ptr1
	**ptr3 = 1111
	fmt.Println("num1:", num1)
	var ptr4 ***int = &ptr3
	***ptr4 = 2222
	fmt.Println("num1:", num1)

	var ptr5 *int

	*ptr5 = 1000

}

func Incr(n int) { // call or pass by value
	n++
}

func IncrP(ptr *int) { // call or pass by reference
	if ptr != nil {
		*ptr++
	}
}
