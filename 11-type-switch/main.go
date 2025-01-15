package main

import "fmt"

func main() {
	var any1 any // here value is nil and type is also nil

	// num, ok := any1.(int)
	// if ok {
	// 	println(num * num)
	// }

	// num1, ok := any1.(int64)
	// if ok {
	// 	println(num1 * num1)
	// }

	// num2, ok := any1.(float32)
	// if ok {
	// 	println(num2 * num2)
	// }

	// num3, ok := any1.(float64)
	// if ok {
	// 	println(num3 * num3)
	// }

	any1 = float32(12.2)
	//any1 = "Hello World"
	println(isNumber(any1))
	switch any1.(type) { // type switch, which can be used for any
	case int:
		fmt.Printf("any1 is of type %T, and square of num: %v\n", any1, any1.(int)*any1.(int))
	case uint:
		fmt.Printf("any1 is of type %T, and square of num: %v\n", any1, any1.(uint)*any1.(uint))
	case uint8:
		fmt.Printf("any1 is of type %T, and square of num: %v\n", any1, uint16(any1.(uint8))*uint16(any1.(uint8)))
	case uint16:
		fmt.Printf("any1 is of type %T, and square of num: %v\n", any1, any1.(uint16)*any1.(uint16))
	case uint32:
		fmt.Printf("any1 is of type %T, and square of num: %v\n", any1, any1.(uint32)*any1.(uint32))
	case uint64:
		fmt.Printf("any1 is of type %T, and square of num: %v\n", any1, any1.(uint64)*any1.(uint64))
	case int8:
		fmt.Printf("any1 is of type %T, and square of num: %v\n", any1, int16(any1.(int8))*int16(any1.(int8)))
	case int16:
		fmt.Printf("any1 is of type %T, and square of num: %v\n", any1, any1.(int16)*any1.(int16))
	case int32:
		fmt.Printf("any1 is of type %T, and square of num: %v\n", any1, any1.(int32)*any1.(int32))
	case int64:
		fmt.Printf("any1 is of type %T, and square of num: %v\n", any1, any1.(int64)*any1.(int64))
	case float32:
		fmt.Printf("any1 is of type %T, and square of num: %v\n", any1, any1.(float32)*any1.(float32))
	case float64:
		fmt.Printf("any1 is of type %T, and square of num: %v\n", any1, any1.(float64)*any1.(float64))
	default:
		fmt.Printf("any1 is of type %T, cannot perform arithmetic operation\n", any1)

	}
}

func isNumber(any1 any) bool {
	switch any1.(type) {
	case int, uint, int8, int16, int32, int64, uint8, uint16, uint32, uint64, float32, float64:
		return true
	default:
		return false
	}
}
