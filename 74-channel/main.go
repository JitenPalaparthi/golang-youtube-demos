package main

func main() {
	ch := make(chan int, 2) // Buffered
	ch <- 100
	ch <- 200
	v := <-ch
	ch <- 300
	println(v)
	v = <-ch
	println(v)
	v = <-ch
	println(v)

}

//					          3 ltr
//Water Head tank ----------------------------------- |Container
