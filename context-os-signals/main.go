package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())

	sigCh := make(chan os.Signal)
	defer close(sigCh)
	signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)

	go func() {
		for {
			select {
			case <-sigCh:
				cancel()
			}
		}
	}()

	ch := GetSq(ctx)

	for sq := range ch {
		fmt.Println(sq)
	}

}

func GetSq(ctx context.Context) chan int {
	ch := make(chan int)
	i := 1
	go func() {
		for {
			select {
			case ch <- i * i:
				i++
				time.Sleep(time.Millisecond * 300)
			case <-ctx.Done():
				fmt.Println("Closing the application")
				time.Sleep(time.Second * 5)
				fmt.Println("Gracefully closing the channel")
				fmt.Println(ctx.Err().Error())
				close(ch)
				return
			}

		}

	}()

	return ch
}

// 1- Generate Squqres of numbers . The sequence starts form 1 and untile it gets cancelled
// 2- I purposely terminate the process , it should inform the GetSq go routine that it received a Os signal
