package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*1)

	//defer cancel()

	// for val := range GenSq(ctx) {
	// 	fmt.Println(val)
	// }
	ch := GenSq(ctx)

	go func(context.Context, chan int) {
		for {
			select {
			case <-ctx.Done():
				cancel()
				fmt.Println("Receiver goroutine:", ctx.Err())
				return
			case val := <-ch:
				fmt.Println(val)
			}
		}
	}(ctx, ch)

	runtime.Goexit() // We know that runtime.Goexit gives a non zero os exit. That means stderr
}

func GenSq(ctx context.Context) chan int {
	ch := make(chan int)
	go func() {
		i := 1
		for {
			select {
			case <-ctx.Done():
				close(ch)
				return
			default:
				ch <- i * i
				i += 1
			}
		}
	}()
	return ch
}
