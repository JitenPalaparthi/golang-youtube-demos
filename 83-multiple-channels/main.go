package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch1 := Generator(5, "Gen-1", time.Millisecond*100)
	ch2 := Generator(3, "Gen-2", time.Millisecond*100)
	ch3 := Generator(4, "Gen-3", time.Millisecond*100)
	ch4 := Generator(4, "Gen-4", time.Millisecond*100)
	sig := FanIn(ch1, ch2, ch3, ch4)
	<-sig
}

func Generator(n uint, name string, d time.Duration) chan string {
	ch := make(chan string)
	go func() {
		for i := 1; i <= int(n); i++ {
			time.Sleep(d)
			ch <- fmt.Sprint(name, "--->", i*i)
		}
		close(ch)
	}()
	return ch
}

func FanIn(chs ...chan string) <-chan struct{} {
	sig := make(chan struct{})
	wg := new(sync.WaitGroup)

	go func() {
		wg.Wait()
		sig <- struct{}{}
	}()

	ChanCall := func(ch1 chan string) {
		for c := range ch1 {
			println(c)
		}
		wg.Done()
	}

	for _, ch := range chs {
		wg.Add(1)
		go ChanCall(ch)
	}

	return sig
}
