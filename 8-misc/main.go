package main

import (
	"errors"
	"fmt"
	"strconv"
	"sync"
)

var (
	WG       *sync.WaitGroup
	MyGlobal int
)

func main() {
	a, b := 10, 20
	// t := a
	// a = b
	// b = t
	a, b = b, a // swapping in golang

	a, b, c := 10, 20, 30

	a, b, c = b, c, a // can also swap three variables can do even more ..

	str1 := "1223"
	//num := int(str1) // cannot type cast from string to int
	num, err := strconv.Atoi(str1)
	if err != nil {
		println(err.Error())
	} else {
		println(num)
	}

	str1 = "12312.23"
	float1, err := strconv.ParseFloat(str1, 32)
	if err != nil {
		println(err.Error())
	} else {
		fmt.Printf("%.2f\n", float1)
	}

	str1 = "k"

	ok, err := strconv.ParseBool(str1)
	if err != nil {
		println(err.Error())
	} else {
		println(ok)
	}

	// convert any value to string

	ok = true
	num1 := 123.123
	hey := "Hello World"
	age := 39

	str2 := fmt.Sprint(hey, " ", age, " ", num1, " ", ok)
	fmt.Printf("Type of str2:%T str2:%s\n", str2, str2)
	str3 := strconv.Itoa(age)
	fmt.Printf("Type of str3:%T str3:%s\n", str3, str3)
	isNumber(100)
}

// type cast
// type assert
// type convert

func isNumber(any1 any) (bool, error) {
	switch any1.(type) {
	case int, uint, int8, int16, int32, int64, uint8, uint16, uint32, uint64, float32, float64:
		return true, nil
	default:
		return false, errors.New("input is nil")
	}
}

func Calc(a, b int) (int, int, int, int) {
	return a + b, a - b, a * b, a / b
}
