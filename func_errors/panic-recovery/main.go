package main

import "fmt"

func mightPanic(shouldPanic bool) {
	if shouldPanic {
		panic("Something wrong")
	}
	fmt.Println("Without panic")
}

func recoverable() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic: ", r)
		}
	}()

	mightPanic(true)
}

func main() {
	recoverable()
}
