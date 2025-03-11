package main

import (
	"sync"
	"time"
)

var (
	wg *sync.WaitGroup = new(sync.WaitGroup)
)

func main() {

	ch1 := Generator(20, time.Millisecond*100) // Does not look like a goroutine
	ch2 := Generator(20, time.Millisecond*100)
	ch3 := Generator(20, time.Millisecond*30)

	wg.Add(4)
	go Receive(wg, ch1, "1")
	go Receive(wg, ch1, "2")
	go Receive(wg, ch3, "3")
	go Receive(wg, ch2, "4")
	wg.Wait()
	println("finished receing all the values, hance the main exits")
	// The gorouine is blocked until it receives a value
}

func Generator(n uint, d time.Duration) <-chan int {
	ch := make(chan int)
	go func() {
		for i := 1; i <= int(n); i++ {
			time.Sleep(d)
			ch <- i * i
		}
		close(ch)
	}()
	return ch
}

// func Receive(n uint, ch <-chan int, sig chan<- bool) {
func Receive(wg *sync.WaitGroup, ch <-chan int, name string) {
	for v := range ch { // The range loop on channels returns only one value, unlike slice and map
		println("reciver->", name, "--", v)
	}
	// how long the range loop iterates, as long as the channel is not closed, range loop tries to fetch values.
	// the moment the channel is closed , range loop exits
	// Finished receiving all values from ch.
	//sig <- true
	println("done with receiving receiver-->", name)
	//sig <- struct{}{}
	wg.Done()
}
