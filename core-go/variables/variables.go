package main

import "fmt"

func main() {
	var greeting string // zero-value is empty string ""
	greeting = "hello, world"

	fmt.Println(greeting)

	var count int
	count = 10
	fmt.Println(count)

	var isRunning bool
	isRunning = true
	fmt.Println(isRunning)

	var firstName, lastName string
	firstName = "Veet"
	lastName = "Moradiya"
	fmt.Println(firstName + " " + lastName)

	email := "test@test.com"
	fmt.Println(email)

	age := 24
	fmt.Println(age)

	var year = 2025
	fmt.Println(year)
}
