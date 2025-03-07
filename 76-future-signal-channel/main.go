package main

import (
	"time"
)

func main() {
	ch := make(chan int)
	// sig := make(chan bool) // it has 1 byte
	sig := make(chan struct{}) // empty struct has zero size

	go Send(10, ch, time.Millisecond*100)
	go Receive(10, ch, sig)
	<-sig
	println("finished receing all the values, hance the main exits")
	// The gorouine is blocked until it receives a value
}

func Send(n uint, ch chan<- int, d time.Duration) {
	for i := 1; i <= int(n); i++ {
		time.Sleep(d)
		ch <- i * i
	}
}

// func Receive(n uint, ch <-chan int, sig chan<- bool) {
func Receive(n uint, ch <-chan int, sig chan<- struct{}) {

	for i := 1; i <= int(n); i++ {
		println(<-ch)
	}
	// Finished receiving all values from ch.
	//sig <- true
	sig <- struct{}{}
}
