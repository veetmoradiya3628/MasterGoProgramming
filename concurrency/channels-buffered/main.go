package main

import "fmt"

func main() {
	messages := make(chan string, 3) // buffered channel
	fmt.Println("Sending messages to buffered channel")
	messages <- "Hello 1"
	messages <- "Hello 2"
	messages <- "Hello 3"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
	fmt.Println(<-messages)

	messages <- "Hello 5"
	messages <- "Hello 6"
	messages <- "Hello 4"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
	fmt.Println(<-messages)

}
