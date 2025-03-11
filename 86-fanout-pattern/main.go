package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	//defer wg.Done()
	for job := range jobs {
		time.Sleep(time.Millisecond * 100) // simulation as the procee takes time .. just assume
		fmt.Println("Worker-->", id, "is processing,", job)
	}
	wg.Done()
}

func main() {
	const (
		numWorkers = 3
		numJobs    = 10
	)

	jobs := make(chan int, numJobs) // buffered channel
	wg := new(sync.WaitGroup)

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, jobs, wg)
	}
	wg.Add(1)
	go func() {
		for j := 1; j <= numJobs; j++ {
			jobs <- j // publisher
		}
		close(jobs)
		wg.Done()
	}()

	wg.Wait()
}
