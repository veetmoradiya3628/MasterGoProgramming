package main

import (
	"fmt"
	"time"
)

type User struct {
	name string
}

func main() {
	messages := make(chan string) // Unbuffered channel
	users := make(chan User)
	go func() {
		fmt.Println("Sending a message to messages channel")
		messages <- "Hello from messages channel"
	}()

	go func() {
		fmt.Println("Sending a message to messages channel")
		messages <- "Hello from messages channel"
	}()

	go func() {
		users <- User{
			name: "Veet",
		}
	}()

	time.Sleep(1 * time.Second)
	fmt.Println("About to get message from channel")
	msg := <-messages
	fmt.Println(msg)

	msg = <-messages
	fmt.Println(msg)

	user := <-users
	fmt.Println("User : ", user)
}
