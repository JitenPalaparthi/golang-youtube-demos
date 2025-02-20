package main

import (
	"fmt"
	"reflect"
)

func main() {

	map1 := make(map[string]any, 5)

	map1["greet"] = func() {
		fmt.Println("Hello World")
	}

	map1["add"] = func(a, b int) int {
		return a + b
	}

	var fn1 Fn = func(i1, i2 int) int {
		return i1 - i2
	}
	map1["sub"] = fn1

	map1["mul"] = Mul

	map1["div"] = Div

	map1["done"] = true
	map1["value"] = 100
	map1["cities"] = []string{"Bangalore", "Trivandrum", "Mumbai", "Vijayawada"}
	a, b := 20, 10
	for key, val := range map1 {

		switch ret := val.(type) {
		case func():
			fmt.Println("Key-->", key)
			fmt.Println("Type of Value:", reflect.TypeOf(val))
			// val.(func())()
			ret()
			println()

		case func(a, b int) int:
			fmt.Println("Key-->", key)
			fmt.Println("Type of Value:", reflect.TypeOf(val))
			r := ret(a, b)
			fmt.Println("result:", r)
			println()

		case Fn:
			fmt.Println("Key-->", key)
			fmt.Println("Type of Value:", reflect.TypeOf(val))
			r := ret(a, b)
			fmt.Println("result:", r)
			println()
		case func(int, int) func() int:
			fmt.Println("Key-->", key)
			fmt.Println("Type of Value:", reflect.TypeOf(val))
			fn := ret(a, b)
			r := fn()
			fmt.Println("result:", r)
			println()

		case []string:
			fmt.Println("Key-->", key)
			fmt.Println("Type of Value:", reflect.TypeOf(val))
			for i, v := range ret {
				fmt.Println("index:", i, "value:", v)
			}
			println()

		case int:
			fmt.Println("Key-->", key)
			fmt.Println("Type of Value:", reflect.TypeOf(val))
			fmt.Println(ret)
			println()

		case bool:
			fmt.Println("Key-->", key)
			fmt.Println("Type of Value:", reflect.TypeOf(val))
			fmt.Println(ret)
			println()

		case string:
			fmt.Println("Key-->", key)
			fmt.Println("Type of Value:", reflect.TypeOf(val))
			fmt.Println(ret)
			println()

		}

	}

}

type Fn func(int, int) int

func Mul(a, b int) int {
	return a * b
}

func Div(a, b int) func() int {
	return func() int {
		return a / b
	}
}
