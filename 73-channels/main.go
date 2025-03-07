package main

import (
	"sync"
	"time"
)

func main() {
	var ch chan int     // This is only the defining a variable but not instantiation
	ch = make(chan int) // instantiation of unbuffered channels
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		//time.Sleep(time.Second * 5)
		ch <- 100 // ch <- sending value or sender
		println("Finish sening value and reveived")
		wg.Done()
	}()

	time.Sleep(time.Second * 5)
	v := <-ch // How long it waits.. Until the sender sends the value
	println(v)
	wg.Wait()
}

// 1. What is a channel
// 		Buffered channel and unbuffered channel
// 2. There is a sender and there is receiver
// 3. Unless the receiver receives the date , sender cannot send the data.
// 4. Unless the sender sends the data, the receiver cannot receive the data.
// 5. There is a block -->
// 		 Sender is sending data.Until the receiver receives the data sender is blocked
//       Recevier is receiving data, Until the sender sends the data, the receiver is blocked
