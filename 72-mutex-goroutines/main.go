package main

import "sync"

var (
	counter int             = 0
	wg      *sync.WaitGroup = new(sync.WaitGroup)
	mu      *sync.Mutex     = new(sync.Mutex)
)

func main() {
	wg.Add(1)
	go func() {
		for i := 1; i <= 1000; i++ {
			wg.Add(1)
			go increment(wg, mu)
		}
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		for i := 1; i <= 1000; i++ {
			wg.Add(1)
			go decrement(wg, mu)
		}
		wg.Done()
	}()
	wg.Wait()
	println(counter)

}

func increment(wg *sync.WaitGroup, mu *sync.Mutex) {
	mu.Lock()
	counter++
	mu.Unlock()
	wg.Done()
}

func decrement(wg *sync.WaitGroup, mu *sync.Mutex) {
	mu.Lock()
	counter--
	mu.Unlock()
	wg.Done()
}
