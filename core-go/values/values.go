package main

import "fmt"

func main() {
	var t any

	fmt.Print("hello world\n")

	fmt.Println("hello " + " world")

	fmt.Println(1 + 1)

	fmt.Println(3.14 + 1.414)

	fmt.Println(true, false)

	fmt.Printf("%+v\n", []int{1, 2, 3}) // slices - array

	fmt.Printf("%+v\n", t)
}
