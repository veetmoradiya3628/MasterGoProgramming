package main

import "fmt"

// This is a simple pipeline example where we have three stages:
// 1. A stage that converts a slice of integers into a channel.
// 2. A stage that squares the integers from the channel.
// 3. A stage that prints the squared integers.

/*
Stage 1: sliceToChannel - This function takes a slice of integers and returns a channel that emits those integers one by one.
It uses a goroutine to send the integers to the channel and closes the channel when done.
*/
func sliceToChannel(nums []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

/*
Stage 2: sq - This function takes a channel of integers and returns a new channel that emits the squares of those integers.
It also uses a goroutine to read from the input channel, compute the square, and send it to the output channel.
*/
func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func main() {
	// input
	nums := []int{1, 3, 4, 2, 5}

	// stage 1
	dataChannel := sliceToChannel(nums)
	// stage 2
	finalChannel := sq(dataChannel)
	// stage 3
	for n := range finalChannel {
		fmt.Print(n, " ")
	}
}
