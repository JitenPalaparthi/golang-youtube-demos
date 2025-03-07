package main

import (
	"sync"
	"time"
)

var (
	wg *sync.WaitGroup = new(sync.WaitGroup)
)

func main() {
	ch := make(chan int)
	defer wg.Wait()

	wg.Add(1)
	go func(*sync.WaitGroup) {
		Send(10, ch, time.Millisecond*100)
		wg.Done()
	}(wg)

	wg.Add(1)
	go func(*sync.WaitGroup) {
		Receive(10, ch)
		wg.Done()
	}(wg)

}

func Send(n uint, ch chan<- int, d time.Duration) {
	for i := 1; i <= int(n); i++ {
		time.Sleep(d)
		ch <- i * i
	}
}

func Receive(n uint, ch <-chan int) {
	for i := 1; i <= int(n); i++ {
		println(<-ch)
	}
}
