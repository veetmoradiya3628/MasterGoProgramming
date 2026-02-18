package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ticker := time.NewTicker(1 * time.Second)
	timer := time.NewTimer(5 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			fmt.Println("Tick")
		case <-timer.C:
			fmt.Println("Timeout")
			return
		}
	}
}

func timerExample() {
	timer := time.NewTimer(1 * time.Second)
	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		<-timer.C
		fmt.Println("After 1 second")

	}()

	for i := 0; i < 10; i++ {
		fmt.Println("Inside main go routine: ", i)
	}

	wg.Wait()
	fmt.Println("Program ends")
}
