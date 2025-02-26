package main

func main() {

	str1 := "Hello World"

	//func() {
	for _, v := range str1 {
		defer print(string(v), " ")
	}
	//}()

	a, b := 100, 200 // a is a global variable w.r.t func 1.1

	defer func(b1 int) { // func 1.1
		println("value of a,b in defer", a, b1)
	}(b)

	a += 1
	b += 1

	println("value of a and b:", a, b)
}
