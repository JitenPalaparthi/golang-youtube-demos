package main

func main() {
	println("start of main")
	//n := 1
	//println(10 / n)

	var arr [3]int
	// for i := 0; i <= len(arr); i++ {
	// println(arr[i])
	// }
	var ptr *int // pointer

	func() { // func 1.1
		println("Hello World")
		func() { // func 1.1.1
			*ptr = 100 // nil pointer dereferencing
			for i := 0; i < len(arr); i++ {
				println(arr[i])
			}
			println("end of func 1.1.1")
		}()
		println("end of func 1.1")
	}()
	//println("end of func 1.1")
	println("End of main")
}

// 1. Error
// 2. panic
// 	runtime panic, panic is triggered by the runtime
//  user defined panic, trigger panic by the developer
// 3. Fatal
