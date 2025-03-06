package main

import (
	"runtime"
	"sync"
	"time"
)

var (
	wg *sync.WaitGroup = new(sync.WaitGroup)
)

func main() {
	defer wg.Wait()
	defer println("End of Main")
	println("Start of Main")

	// defer func() {
	// go Greet() // This is run as a concurrent go routine
	// }()
	//wg.Add(3)

	wg.Add(1)
	go func() {
		Greet()   // This is run as a concurrent go routine
		wg.Done() // Done is called to decrement the counter,once it the goroutine is finished
	}()

	wg.Add(1)
	go GreetMe(wg)

	wg.Add(1)
	go func() {
		a, b := 0, 1
		for i := 1; i <= 5; i++ {
			time.Sleep(time.Millisecond * 100)
			println("Fib", a)
			a, b = b, a+b
		}
		wg.Done()
	}()

	wg.Add(1)
	go func(wg *sync.WaitGroup) { // func 1.3
		for i := 1; ; i++ {
			time.Sleep(time.Millisecond * 100)
			if i > 10 {
				wg.Done()
				runtime.Goexit() // who calls it? func 1.3
			}
			if i%2 == 0 {
				println("even number", i)
			}
		}
	}(wg)

	println("Just to check whether this println is executed or not..")
	//wg.Wait()
} // return from main which is os.Exit(0)

func Greet() {
	//time.Sleep(time.Millisecond * 2)
	println("Hello, World!")
}

func GreetMe(wg *sync.WaitGroup) {
	defer wg.Done()
	println("Hello, Me!")
}

// 1. Main is also a goroutine
// 2. No Goroutine waits for other goroutine to complete its execution
// 3. The order of execution of goroutines is not guaranteed
// 4. Can assume it but there is no parent or child goroutine
