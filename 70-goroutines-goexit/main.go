package main

import (
	"runtime"
)

func main() {
	defer println("End of Main")
	println("Start of Main")
	go Greet() // This is run as a concurrent go routine
	go func() {
		a, b := 0, 1
		for i := 1; i <= 5; i++ {
			go println("Fib", a)
			a, b = b, a+b
		}
	}()

	go func() { // func 1.3
		for i := 1; ; i++ {
			if i > 10 {
				runtime.Goexit() // who calls it? func 1.3
			}
			if i%2 == 0 {
				go println("even number", i)
			}
		}
	}()

	// go func() {
	// for {
	//
	// }
	// }()
	//time.Sleep(time.Second * 1)
	runtime.Goexit() // the main is returned not gracefully or where it is supposed to be
	println("Just to check whether this println is executed or not..")
} // return from main which is os.Exit(0)

func Greet() {
	//time.Sleep(time.Millisecond * 2)
	println("Hello, World!")
}

// 1. Main is also a goroutine
// 2. No Goroutine waits for other goroutine to complete its execution
// 3. The order of execution of goroutines is not guaranteed
// 4. Can assume it but there is no parent or child goroutine
