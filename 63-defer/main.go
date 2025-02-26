package main

func main() {
	defer println("End of main-1") // 1 defer
	defer func() {                 // func 1.1 // 3 defer
		defer func() { // func 1.2 // 4 defer
			println("Hello World!")
		}()
		defer println("End of func 1.1") // 5 defer
		println("start of func 1.1")
	}()
	defer println("End of main-2") // 2 defer
	println("start of main")
}

// defer defers the execution to the end of the call stack
// all deferred calls are stored in the stack for a caller stack frame
