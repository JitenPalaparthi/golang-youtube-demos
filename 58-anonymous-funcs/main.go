package main

func main() {

	func() {
		println("Hello World")
	}()

	c1 := func(a, b int) int {
		return a + b
	}(10, 20)
	println(c1)

	f1 := func(a, b int) int {
		return a - b
	} // because there is no executor, it is a function

	c2 := f1(10, 20)
	println(c2)

	var f2 func(int, int) int = Mul
	c3 := f2(10, 20)
	println(c3)
}

func Mul(a, b int) int {
	return a * b
}
