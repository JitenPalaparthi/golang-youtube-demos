package main

import (
	"sync"
	"time"
)

var (
	wg *sync.WaitGroup = new(sync.WaitGroup)
)

func main() {
	ch := make(chan int, 3)
	// sig := make(chan bool) // it has 1 byte
	sig := make(chan struct{}) // empty struct has zero size

	go Send(20, ch, time.Millisecond*0)
	go Receive(ch, sig, "1")
	go Receive(ch, sig, "2")
	go Receive(ch, sig, "3")
	<-sig
	<-sig
	<-sig
	println("finished receing all the values, hance the main exits")
	// The gorouine is blocked until it receives a value
}

func Send(n uint, ch chan<- int, d time.Duration) {
	for i := 1; i <= int(n); i++ {
		time.Sleep(d)
		ch <- i * i
	}
	close(ch)
}

// func Receive(n uint, ch <-chan int, sig chan<- bool) {
func Receive(ch <-chan int, sig chan<- struct{}, name string) {
	for v := range ch { // The range loop on channels returns only one value, unlike slice and map
		println("reciver->", name, "--", v)
	}
	// how long the range loop iterates, as long as the channel is not closed, range loop tries to fetch values.
	// the moment the channel is closed , range loop exits
	// Finished receiving all values from ch.
	//sig <- true
	println("done with receiving receiver-->", name)
	sig <- struct{}{}
}
