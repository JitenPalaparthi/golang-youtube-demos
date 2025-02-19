package main

import "fmt"

func main() {

	r1 := Calc(10, 20, func(i1, i2 int) int { return i1 + i2 })
	fmt.Println("Sum of a and b:", r1)

	f1 := func(a, b int) int {
		return a - b
	}

	r2 := Calc(10, 20, f1)

	fmt.Println("Sub of a and b:", r2)

	r3 := Calc(10, 20, Mul)

	fmt.Println("Mul of a and b:", r3)
	var fn1 Fn = func(i1, i2 int) int {
		return i1 / i2
	}

	fmt.Println(fn1.Sq(100))

	r4 := Calc(10, 20, fn1)
	fmt.Println("Div of a and b:", r4)

	r5 := Calc1(10, 20, (Fn)(Mul))
	fmt.Println("Mul of a and b:", r5)

}

type Fn func(int, int) int

func (f Fn) Sq(a int) int {
	return a * a
}

func Calc(a, b int, fn func(int, int) int) int {
	return fn(a, b)
}

func Calc1(a, b int, fn Fn) int {
	return fn(a, b)
}

func Mul(a, b int) int {
	return a * b
}
