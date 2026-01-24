package main

import "fmt"

func simpleDefer() {
	fmt.Println("Function simpleDefer: Start")
	defer fmt.Println("Function simpleDefer: End")
	fmt.Println("Function simpleDefer: Middle1")
	fmt.Println("Function simpleDefer: Middle2")
	fmt.Println("Function simpleDefer: Middle3")
}

func lifeOfDefer() {
	fmt.Println("Function lifeOfDefer: Start")
	defer fmt.Println("Function lifeOfDefer: Defer1")
	defer fmt.Println("Function lifeOfDefer: Defer2")
	fmt.Println("Function lifeOfDefer: Middle")

}

func main() {

	defer func() {
		fmt.Println("Before the return of main()")
	}()

	// simpleDefer()
	lifeOfDefer()
	fmt.Println("Last in main")
}
