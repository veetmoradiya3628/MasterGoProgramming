package main

import (
	"context"
	"fmt"
	"time"
)

/*
What is context in Go?
- Context is a package in Go that provides a way to carry deadlines, cancellation signals, and other request-scoped values across API boundaries and between processes. It is used to manage the lifecycle of a request and to propagate cancellation signals across goroutines.

How to use context in Go?
- To use context in Go, you can create a new context using the context.Background() function, which returns an empty context. You can then use the context.WithCancel() function to create a new context that can be cancelled. You can also use the context.WithTimeout() function to create a new context that will automatically cancel after a specified duration.

In this example, we have two goroutines, ping and pong, that send messages to a channel every second.
The main function creates a context with a cancel function and starts the ping and pong goroutines.
It also starts another goroutine that listens for messages from the channel and prints them until a timeout occurs after 5 seconds, at which point it cancels the context and signals that the operation is completed.

*/

func ping(ctx context.Context, ch chan string) {
	for {
		select {
		case <-ctx.Done():
			return
		case ch <- fmt.Sprintf("ping: %v", time.Now()):
			time.Sleep(1 * time.Second)
		}
	}
}

func pong(ctx context.Context, ch chan string) {
	for {
		select {
		case <-ctx.Done():
			return
		case ch <- fmt.Sprintf("pong: %v", time.Now()):
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background()) // create a context with cancel function
	defer cancel()

	pingerChannel := make(chan string)
	done := make(chan struct{})

	go ping(ctx, pingerChannel)
	go pong(ctx, pingerChannel)

	go func() {
		timeout := time.After(5 * time.Second)
		for {
			select {
			case <-timeout:
				fmt.Println("Operation completed")
				close(pingerChannel)
				done <- struct{}{}
				return
			case msg := <-pingerChannel:
				fmt.Println(msg)
			}
		}
	}()
	<-done
	fmt.Println("done")
}
