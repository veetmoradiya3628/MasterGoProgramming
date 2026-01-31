package main

import (
	"fmt"
	"time"
)

func sayHello(message string, delay time.Duration) {
	time.Sleep(delay)
	fmt.Println("sayHello", message)
}

// main go routine
func main() {
	fmt.Println("Hello from main() Goroutine")
	go sayHello("Hello World 1", time.Second)
	go sayHello("Hello World 2 one", 2*time.Second)
	go sayHello("Hello World 2 two", 2*time.Second)
	go sayHello("Hello World 3 ", 3*time.Second)
	fmt.Println("Last message from main")
	time.Sleep(2 * time.Second) // problem because it's not efficient way of doing the stuff in dynamic environment
}
