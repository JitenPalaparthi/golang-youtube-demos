package main

import (
	"fmt"
	"time"
)

func main() {

	ch1 := Generator(10, "Gen-1", time.Millisecond*100)
	ch2 := Generator(10, "Gen-2", time.Millisecond*100)
	ch3 := Generator(10, "Gen-3", time.Millisecond*100)

	// wg := new(sync.WaitGroup)
	// wg.Add(1)

	//c := 0
	done1, done2, done3 := false, false, false
	sig := make(chan struct{})

	go func() {
	out:
		for {
			if done1 && done2 && done3 {
				// wg.Done()
				//runtime.Goexit()
				break out
			}
			select { // Select is used to work with channels
			case v, ok := <-ch1:
				if ok {
					println(v)
				} else {
					done1 = true
				}

			case v, ok := <-ch2:
				if ok {
					println(v)
				} else {
					done2 = true
				}

			case v, ok := <-ch3:
				if ok {
					println(v)
				} else {
					done3 = true
				}
			}
		}
		sig <- struct{}{}
	}()

	//wg.Wait()
	<-sig

}

func Generator(n uint, name string, d time.Duration) <-chan string {
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
