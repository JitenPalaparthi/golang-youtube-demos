package main

func main() {

	// for i := 1; i <= 10; i++ {
	// 	println(i)
	// }

	println("fib numbers")
	a, b := 0, 1

	for i := 1; i <= 50; i++ {
		print(a, " ")
		a, b = b, a+b // swapping
	}
	println()

	println("fib of even")
	a, b = 0, 1

	for i := 1; i <= 50; i++ {
		if a%2 == 0 {
			print(a, " ")
		}
		a, b = b, a+b // swapping
	}
	println()

	println("odd numbers")
	for i := 1; i <= 50; i++ {
		if i%2 == 1 {
			continue
		}
		print(i, " ")
	}
	println()

	// for loop like a while loop

	println("multiples of 3")
	i := 1
	for i <= 50 { // for loop but like a while loop
		if i%3 == 0 {
			print(i, " ")
		}
		i++
	}
	println()
	println("multiples of 4")

	for i := 1; ; i++ {
		if i == 50 {
			break
		}
		if i%4 == 0 {
			print(i, " ")
		}
	}
	println()

	println("multiples of 5")

	for i := 1; ; {
		if i == 50 {
			break
		}
		if i%5 == 0 {
			print(i, " ")
		}
		i++
	}
	println()

	println("unconditional loop")
	println("multiples of 6")
	i = 1
	for { // loop is unconditional but use return to return from the loop
		if i == 50 {
			return // returns from the caller(main here)
		}
		if i%6 == 0 {
			print(i, " ")
		}
		i++
	}
	println()
}
