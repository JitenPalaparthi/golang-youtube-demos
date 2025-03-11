package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	worker := func(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
		defer wg.Done()
		for job := range jobs {
			time.Sleep(time.Millisecond * 100)
			fmt.Println("Worker-->", id, "is processing,", job)
			results <- job * job
		}
	}

	const (
		numWorkers = 3
		numJobs    = 10
	)

	jobs := make(chan int, numJobs)    // buffered channel
	results := make(chan int, numJobs) // buffered channel
	wg := new(sync.WaitGroup)

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, jobs, results, wg)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j // publisher
	}
	close(jobs)

	wg.Wait()
	close(results) // even it is closed

	for result := range results { // even the channel is closed, all the values in the channel can be retrived
		fmt.Println("Result:", result)
	}

}
