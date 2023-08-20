package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	//ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*1)
	//ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Millisecond*1))
	ctx, cancel := context.WithCancel(context.Background())

	chtime := time.After(time.Millisecond * 1)

	go func() {
	outer:
		for {
			select {
			case <-chtime:
				fmt.Println("Context cancelled")
				cancel()
				break outer
			}
		}
	}()

	ch := GenSq(ctx)

	for val := range ch {
		fmt.Println(val)
	}

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
				i++
			}

		}
	}()

	return ch
}
