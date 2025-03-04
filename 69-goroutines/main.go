package main

import "time"

func main() {
	defer println("End of Main")
	println("Start of Main")
	go Greet() // This is run as a concurrent go routine
	go func() {
		a, b := 0, 1
		for i := 1; i <= 100; i++ {
			go println("Fib", a)
			a, b = b, a+b
		}
	}()

	go func() {
		for i := 1; i <= 100; i++ {
			if i%2 == 0 {
				go println("even number", i)
			}
		}
	}()
	time.Sleep(time.Millisecond * 100) // This is not a solution
} // return from main which is os.Exit(0)

func Greet() {
	//time.Sleep(time.Millisecond * 2)
	println("Hello, World!")
}

// 1. Main is also a goroutine
// 2. No Goroutine waits for other goroutine to complete its execution
// 3. The order of execution of goroutines is not guaranteed
// 4. Can assume it but there is no parent or child goroutine
